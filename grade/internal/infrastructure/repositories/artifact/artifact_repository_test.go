package artifact

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/artifact"
	artifactVal "github.com/emacsway/grade/grade/internal/domain/artifact/values"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/seedwork/domain/aggregate"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session/pgx"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/utils/testutils"
)

type testCase func(t *testing.T, repositoryOption RepositoryOption)

func TestArtifactRepository(t *testing.T) {

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
			r, err := repositoryOption.Session.Exec("DELETE FROM artifact")
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
	var exporterActual artifact.ArtifactExporter
	factory := artifact.NewArtifactFaker(
		artifact.WithTenantId(repositoryOption.TenantId),
		artifact.WithTransientId(),
	)
	factory.OwnerId.MemberId = repositoryOption.MemberId
	agg, err := factory.Create()
	require.NoError(t, err)
	err = repositoryOption.Repository.Insert(agg, aggregate.EventMeta{})
	require.NoError(t, err)
	agg.Export(&exporterActual)
	assert.Greater(t, int(exporterActual.Id.ArtifactId), 0)
}

func testGet(t *testing.T, repositoryOption RepositoryOption) {
	var exporterActual artifact.ArtifactExporter
	var exporterExpected artifact.ArtifactExporter
	factory := NewArtifactFaker(
		repositoryOption.Session,
		artifact.WithTenantId(repositoryOption.TenantId),
	)
	err := factory.BuildDependencies()
	require.NoError(t, err)
	aggExpected, err := factory.Create()
	require.NoError(t, err)
	aggExpected.Export(&exporterExpected)

	assert.Greater(t, int(exporterExpected.Id.ArtifactId), 0)

	id, err := artifactVal.NewArtifactId(
		uint(exporterExpected.Id.TenantId),
		uint(exporterExpected.Id.ArtifactId),
	)
	require.NoError(t, err)
	aggActual, err := repositoryOption.Repository.Get(id)
	require.NoError(t, err)
	aggActual.Export(&exporterActual)
	assert.Equal(t, exporterExpected, exporterActual)
}

type RepositoryOption struct {
	Name       string
	Repository *ArtifactRepository
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
	currentSession := pgx.NewPgxSession(db)
	tf := tenantRepo.NewTenantFaker(currentSession)
	aTenant, err := tf.Create()
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
		Repository: NewArtifactRepository(currentSession),
		Session:    currentSession,
		TenantId:   uint(tenantExp.Id),
		MemberId:   uint(memberFaker.Id.MemberId),
	}
}
