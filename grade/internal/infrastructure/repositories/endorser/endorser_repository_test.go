package endorser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/testutils"
)

type testCase func(t *testing.T, repositoryOption RepositoryOption)

func TestEndorserRepository(t *testing.T) {

	repositoryOptions := createRepositories(t)

	for i := range repositoryOptions {
		ro := repositoryOptions[i]

		t.Run(ro.Name, func(t *testing.T) {
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
		/* TODO:
			defer func() {
			r, err := repositoryOption.Session.Exec("DELETE FROM endorser")
			require.NoError(t, err)
			rowsAffected, err := r.RowsAffected()
			require.NoError(t, err)
			assert.Greater(t, int(rowsAffected), 0)
		}()
		*/
		callable(t, repositoryOption)
	}
}

func testInsert(t *testing.T, repositoryOption RepositoryOption) {
	var exporterExpected endorser.EndorserExporter
	var exporterActual endorser.EndorserExporter
	factory := endorser.NewEndorserFaker(
		endorser.WithTenantId(repositoryOption.TenantId),
		endorser.WithMemberId(repositoryOption.MemberId),
	)
	agg, err := factory.Create()
	require.NoError(t, err)
	err = repositoryOption.Repository.Insert(agg)
	require.NoError(t, err)
	agg.Export(&exporterExpected)

	id, err := memberVal.NewMemberId(
		uint(exporterExpected.Id.TenantId),
		uint(exporterExpected.Id.MemberId),
	)
	require.NoError(t, err)
	aggRead, err := repositoryOption.Repository.Get(id)
	require.NoError(t, err)
	aggRead.Export(&exporterActual)
	assert.Equal(t, exporterExpected, exporterActual)
}

func testGet(t *testing.T, repositoryOption RepositoryOption) {
	var exporterExpected endorser.EndorserExporter
	var exporterActual endorser.EndorserExporter
	factory := NewEndorserFaker(
		repositoryOption.Session,
	)
	err := factory.BuildDependencies()
	require.NoError(t, err)
	aggExpected, err := factory.Create()
	require.NoError(t, err)
	aggExpected.Export(&exporterExpected)
	assert.Greater(t, int(exporterExpected.Id.MemberId), 0)

	id, err := memberVal.NewMemberId(
		uint(exporterExpected.Id.TenantId),
		uint(exporterExpected.Id.MemberId),
	)
	require.NoError(t, err)
	aggActual, err := repositoryOption.Repository.Get(id)
	require.NoError(t, err)
	aggActual.Export(&exporterActual)
	assert.Equal(t, exporterExpected, exporterActual)
}

type RepositoryOption struct {
	Name       string
	Repository *EndorserRepository
	Session    session.DbSession
	TenantId   uint
	MemberId   uint
}

func createRepositories(t *testing.T) []RepositoryOption {
	return []RepositoryOption{
		newPostgresqlRepositoryOption(t),
	}
}

func newPostgresqlRepositoryOption(t *testing.T) RepositoryOption {
	var tenantExp tenant.TenantExporter
	db, err := testutils.NewTestDb()
	require.NoError(t, err)
	currentSession := session.NewPgxSession(db)
	tenantFaker := tenantRepo.NewTenantFaker(currentSession)
	aTenant, err := tenantFaker.Create()
	require.NoError(t, err)
	aTenant.Export(&tenantExp)
	memberFaker := memberRepo.NewMemberFaker(
		currentSession,
		member.WithTenantId(uint(tenantExp.Id)),
	)
	_, err = memberFaker.Create()
	require.NoError(t, err)
	return RepositoryOption{
		Name:       "PostgreSQL",
		Repository: NewEndorserRepository(currentSession),
		Session:    currentSession,
		TenantId:   uint(tenantExp.Id),
		MemberId:   uint(memberFaker.Id.MemberId),
	}
}
