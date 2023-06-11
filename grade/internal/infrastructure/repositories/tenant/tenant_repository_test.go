package tenant

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/seedwork"
)

type testCase func(t *testing.T, repositoryOption RepositoryOption)

func TestTenantRepository(t *testing.T) {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	repositoryOption := createRepositories(t)

	for i := range repositoryOption {
		// When you are looping over slice and later using iterated value in goroutine (here because of t.Parallel()),
		// you need to always create variable scoped in loop body!
		// More info here: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		r := repositoryOption[i]

		t.Run(r.Name, func(t *testing.T) {
			// It's always a good idea to build all non-unit tests to be able to work in parallel.
			// Thanks to that, your tests will be always fast and you will not be afraid to add more tests because of slowdown.
			t.Parallel()

			t.Run("testInsert", func(t *testing.T) {
				t.Parallel()
				clearable(testInsert)(t, r)
			})
		})
	}

}

func clearable(callable testCase) testCase {
	return func(t *testing.T, repositoryOption RepositoryOption) {
		callable(t, repositoryOption)

		r, err := repositoryOption.Session.Exec("DELETE FROM tenant")
		require.NoError(t, err)
		rowsAffected, err := r.RowsAffected()
		require.NoError(t, err)
		assert.Greater(t, int(rowsAffected), 0)
	}
}

func testInsert(t *testing.T, repositoryOption RepositoryOption) {
	var actualExporter tenant.TenantExporter
	factory := tenant.NewTenantFakeFactory(tenant.WithTransientId())
	agg, err := factory.Create()
	require.NoError(t, err)
	repositoryOption.Repository.Insert(agg)
	agg.Export(&actualExporter)
	assert.Greater(t, int(actualExporter.Id), 0)
}

type RepositoryOption struct {
	Name       string
	Repository TenantRepository
	Session    infrastructure.DbSessionExecutor
}

func createRepositories(t *testing.T) []RepositoryOption {
	return []RepositoryOption{
		newPostgresqlRepositoryOption(t),
	}
}

func newPostgresqlRepositoryOption(t *testing.T) RepositoryOption {
	db, err := seedwork.NewTestDb()
	session := infrastructure.NewPgxSession(db)
	require.NoError(t, err)
	return RepositoryOption{
		Name:       "PostgreSQL",
		Repository: NewTenantRepository(session),
		Session:    session,
	}
}