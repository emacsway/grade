package tenant

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/tenant"
	tenantVal "github.com/emacsway/grade/grade/internal/domain/tenant/values"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/session"
	"github.com/emacsway/grade/grade/internal/seedwork/infrastructure/testutils"
)

type testCase func(t *testing.T, repositoryOption RepositoryOption)

func TestTenantRepository(t *testing.T) {
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
			r, err := repositoryOption.Session.Exec("DELETE FROM tenant")
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
	var exporterActual tenant.TenantExporter
	factory := tenant.NewTenantFaker(tenant.WithTransientId())
	agg, err := factory.Create()
	require.NoError(t, err)
	err = repositoryOption.Repository.Insert(agg)
	require.NoError(t, err)
	agg.Export(&exporterActual)
	assert.Greater(t, int(exporterActual.Id), 0)
}

func testGet(t *testing.T, repositoryOption RepositoryOption) {
	var exporterExpected tenant.TenantExporter
	var exporterActual tenant.TenantExporter
	factory := NewTenantFaker(repositoryOption.Session)
	aggExpected, errActual := factory.Create()
	require.NoError(t, errActual)
	aggExpected.Export(&exporterExpected)
	assert.Greater(t, int(exporterExpected.Id), 0)

	id, errActual := tenantVal.NewTenantId(uint(exporterExpected.Id))
	require.NoError(t, errActual)
	aggRead, errActual := repositoryOption.Repository.Get(id)
	require.NoError(t, errActual)
	aggRead.Export(&exporterActual)
	assert.Equal(t, exporterExpected, exporterActual)
}

type RepositoryOption struct {
	Name       string
	Repository *TenantRepository
	Session    session.DbSession
}

func createRepositories(t *testing.T) []RepositoryOption {
	return []RepositoryOption{
		newPostgresqlRepositoryOption(t),
	}
}

func newPostgresqlRepositoryOption(t *testing.T) RepositoryOption {
	db, err := testutils.NewTestDb()
	require.NoError(t, err)
	currentSession := session.NewPgxSession(db)
	return RepositoryOption{
		Name:       "PostgreSQL",
		Repository: NewTenantRepository(currentSession),
		Session:    currentSession,
	}
}
