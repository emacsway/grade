package member

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/seedwork"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
)

type testCase func(t *testing.T, repositoryOption RepositoryOption)

func TestMemberRepository(t *testing.T) {
	// TODO: r := rand.New(rand.NewSource(time.Now().UnixNano()))

	repositoryOptions := createRepositories(t)

	for i := range repositoryOptions {
		// When you are looping over slice and later using iterated value in goroutine (here because of t.Parallel()),
		// you need to always create variable scoped in loop body!
		// More info here: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		ro := repositoryOptions[i]

		t.Run(ro.Name, func(t *testing.T) {
			// It's always a good idea to build all non-unit tests to be able to work in parallel.
			// Thanks to that, your tests will be always fast and you will not be afraid to add more tests because of slowdown.
			t.Parallel()

			t.Run("testInsert", func(t *testing.T) {
				t.Parallel()
				clearable(testInsert)(t, ro)
			})

			t.Run("testGet", func(t *testing.T) {
				t.Parallel()
				clearable(testGet)(t, ro)
			})
		})
	}
}

func clearable(callable testCase) testCase {
	return func(t *testing.T, repositoryOption RepositoryOption) {
		/* defer func() {
			r, err := repositoryOption.Session.Exec("DELETE FROM member")
			require.NoError(t, err)
			rowsAffected, err := r.RowsAffected()
			require.NoError(t, err)
			assert.Greater(t, int(rowsAffected), 0)
		}() */
		callable(t, repositoryOption)
	}
}

func testInsert(t *testing.T, repositoryOption RepositoryOption) {
	var actualExporter member.MemberExporter
	factory := member.NewMemberFakeFactory(
		member.WithTenantId(repositoryOption.TenantId),
		member.WithTransientId(),
	)
	agg, err := factory.Create()
	require.NoError(t, err)
	err = repositoryOption.Repository.Insert(agg)
	require.NoError(t, err)
	agg.Export(&actualExporter)
	assert.Greater(t, int(actualExporter.Id.MemberId), 0)
}

func testGet(t *testing.T, repositoryOption RepositoryOption) {
	var exporterActual member.MemberExporter
	var exporterRead member.MemberExporter
	factory := NewMemberFakeFactory(
		repositoryOption.Session,
		member.WithTenantId(repositoryOption.TenantId),
	)
	agg, err := factory.Create()
	require.NoError(t, err)
	agg.Export(&exporterActual)
	assert.Greater(t, int(exporterActual.Id.MemberId), 0)

	id, err := memberVal.NewTenantMemberId(
		uint(exporterActual.Id.TenantId),
		uint(exporterActual.Id.MemberId),
	)
	require.NoError(t, err)
	aggRead, err := repositoryOption.Repository.Get(id)
	require.NoError(t, err)
	aggRead.Export(&exporterRead)
	assert.Equal(t, exporterActual, exporterRead)
}

type RepositoryOption struct {
	Name       string
	Repository *MemberRepository
	Session    infrastructure.DbSession
	TenantId   uint
}

func createRepositories(t *testing.T) []RepositoryOption {
	return []RepositoryOption{
		newPostgresqlRepositoryOption(t),
	}
}

func newPostgresqlRepositoryOption(t *testing.T) RepositoryOption {
	var tenantExp tenant.TenantExporter
	db, err := seedwork.NewTestDb()
	require.NoError(t, err)
	session := infrastructure.NewPgxSession(db)
	tf := tenantRepo.NewTenantFakeFactory(session)
	tenant, err := tf.Create()
	require.NoError(t, err)
	tenant.Export(&tenantExp)
	return RepositoryOption{
		Name:       "PostgreSQL",
		Repository: NewMemberRepository(session),
		Session:    session,
		TenantId:   uint(tenantExp.Id),
	}
}