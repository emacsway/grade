package competence

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/competence"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/session"
	"github.com/emacsway/grade/grade/internal/infrastructure/seedwork/testutils"
)

type testCase func(t *testing.T, repositoryOption RepositoryOption)

func TestCompetenceRepository(t *testing.T) {

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
			r, err := repositoryOption.Session.Exec("DELETE FROM competence")
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
	var exporterActual competence.CompetenceExporter
	factory := competence.NewCompetenceFaker(
		competence.WithTenantId(repositoryOption.TenantId),
		competence.WithTransientId(),
	)
	factory.OwnerId.MemberId = repositoryOption.MemberId
	agg, err := factory.Create()
	require.NoError(t, err)
	err = repositoryOption.Repository.Insert(agg)
	require.NoError(t, err)
	agg.Export(&exporterActual)
	assert.Greater(t, int(exporterActual.Id.CompetenceId), 0)
}

func testGet(t *testing.T, repositoryOption RepositoryOption) {
	var exporterActual competence.CompetenceExporter
	var exporterRead competence.CompetenceExporter
	factory := NewCompetenceFaker(
		repositoryOption.Session,
		competence.WithTenantId(repositoryOption.TenantId),
	)
	factory.OwnerId.MemberId = repositoryOption.MemberId
	agg, err := factory.Create()
	require.NoError(t, err)
	agg.Export(&exporterActual)
	assert.Greater(t, int(exporterActual.Id.CompetenceId), 0)

	id, err := competenceVal.NewCompetenceId(
		uint(exporterActual.Id.TenantId),
		uint(exporterActual.Id.CompetenceId),
	)
	require.NoError(t, err)
	aggRead, err := repositoryOption.Repository.Get(id)
	require.NoError(t, err)
	aggRead.Export(&exporterRead)
	assert.Equal(t, exporterActual, exporterRead)
}

type RepositoryOption struct {
	Name       string
	Repository *CompetenceRepository
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
	session := session.NewPgxSession(db)
	tf := tenantRepo.NewTenantFaker(session)
	aTenant, err := tf.Create()
	require.NoError(t, err)
	aTenant.Export(&tenantExp)
	memberFaker := memberRepo.NewMemberFaker(
		session,
		member.WithTenantId(uint(tenantExp.Id)),
	)
	_, err = memberFaker.Create()
	require.NoError(t, err)
	return RepositoryOption{
		Name:       "PostgreSQL",
		Repository: NewCompetenceRepository(session),
		Session:    session,
		TenantId:   uint(tenantExp.Id),
		MemberId:   uint(memberFaker.Id.MemberId),
	}
}
