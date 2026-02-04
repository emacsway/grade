package member

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/utils/testutils"
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
		/* TODO:
		defer func() {
			err := repositoryOption.Pool.Session(context.Background(), func(s session.Session) error {
				_, err := s.(session.DbSession).Connection().Exec("TRUNCATE member CASCADE")
				return err
			})
			if err != nil {
				t.Logf("cleanup warning: %v", err)
			}
		}()
		*/
		callable(t, repositoryOption)
	}
}

func testInsert(t *testing.T, repositoryOption RepositoryOption) {
	err := repositoryOption.Pool.Session(context.Background(), func(s session.Session) error {
		var exporterActual member.MemberExporter
		factory := member.NewMemberFaker(
			member.WithTenantId(repositoryOption.TenantId),
			member.WithTransientId(),
		)
		agg, err := factory.Create(s)
		require.NoError(t, err)
		err = repositoryOption.Repository.Insert(s, agg)
		require.NoError(t, err)
		agg.Export(&exporterActual)
		assert.Greater(t, int(exporterActual.Id.MemberId), 0)
		return nil
	})
	require.NoError(t, err)
}

func testGet(t *testing.T, repositoryOption RepositoryOption) {
	err := repositoryOption.Pool.Session(context.Background(), func(s session.Session) error {
		var exporterExpected member.MemberExporter
		var exporterActual member.MemberExporter
		factory := NewMemberFaker(member.WithTenantId(repositoryOption.TenantId))
		aggExpected, err := factory.Create(s)
		require.NoError(t, err)
		aggExpected.Export(&exporterExpected)
		assert.Greater(t, int(exporterExpected.Id.MemberId), 0)

		id, err := memberVal.NewMemberId(
			uint(exporterExpected.Id.TenantId),
			uint(exporterExpected.Id.MemberId),
		)
		require.NoError(t, err)
		aggActual, err := repositoryOption.Repository.Get(s, id)
		require.NoError(t, err)
		aggActual.Export(&exporterActual)
		assert.Equal(t, exporterExpected, exporterActual)
		return nil
	})
	require.NoError(t, err)
}

type RepositoryOption struct {
	Name       string
	Repository *MemberRepository
	Pool       session.SessionPool
	TenantId   uint
}

func createRepositories(t *testing.T) []RepositoryOption {
	return []RepositoryOption{
		newPostgresqlRepositoryOption(t),
	}
}

func newPostgresqlRepositoryOption(t *testing.T) RepositoryOption {
	var tenantExp tenant.TenantExporter
	pool, err := testutils.NewPgSessionPool()
	require.NoError(t, err)

	var tenantId uint
	err = pool.Session(context.Background(), func(s session.Session) error {
		tf := tenantRepo.NewTenantFaker()
		aTenant, err := tf.Create(s)
		require.NoError(t, err)
		aTenant.Export(&tenantExp)
		tenantId = uint(tenantExp.Id)
		return nil
	})
	require.NoError(t, err)

	return RepositoryOption{
		Name:       "PostgreSQL",
		Repository: NewMemberRepository(),
		Pool:       pool,
		TenantId:   tenantId,
	}
}
