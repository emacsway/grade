package competence

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/competence"
	competenceVal "github.com/emacsway/grade/grade/internal/domain/competence/values"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/utils/testutils"
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
			err := repositoryOption.Pool.Session(context.Background(), func(s session.Session) error {
				_, err := s.(session.DbSession).Connection().Exec("TRUNCATE competence CASCADE")
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
		var exporterActual competence.CompetenceExporter
		factory := competence.NewCompetenceFaker(
			competence.WithTenantId(repositoryOption.TenantId),
			competence.WithTransientId(),
		)
		factory.OwnerId.MemberId = repositoryOption.MemberId
		agg, err := factory.Create(s)
		require.NoError(t, err)
		err = repositoryOption.Repository.Insert(s, agg)
		require.NoError(t, err)
		agg.Export(&exporterActual)
		assert.Greater(t, int(exporterActual.Id.CompetenceId), 0)
		return nil
	})
	require.NoError(t, err)
}

func testGet(t *testing.T, repositoryOption RepositoryOption) {
	err := repositoryOption.Pool.Session(context.Background(), func(s session.Session) error {
		var exporterExpected competence.CompetenceExporter
		var exporterActual competence.CompetenceExporter
		factory := NewCompetenceFaker(competence.WithTenantId(repositoryOption.TenantId))
		factory.OwnerId.MemberId = repositoryOption.MemberId
		aggExpected, err := factory.Create(s)
		require.NoError(t, err)
		aggExpected.Export(&exporterExpected)
		assert.Greater(t, int(exporterExpected.Id.CompetenceId), 0)

		id, err := competenceVal.NewCompetenceId(
			uint(exporterExpected.Id.TenantId),
			uint(exporterExpected.Id.CompetenceId),
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
	Repository *CompetenceRepository
	Pool       session.SessionPool
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
	pool, err := testutils.NewPgSessionPool()
	require.NoError(t, err)

	var tenantId uint
	var memberId uint
	err = pool.Session(context.Background(), func(s session.Session) error {
		tf := tenantRepo.NewTenantFaker()
		aTenant, err := tf.Create(s)
		require.NoError(t, err)
		aTenant.Export(&tenantExp)
		tenantId = uint(tenantExp.Id)

		memberFaker := memberRepo.NewMemberFaker(member.WithTenantId(tenantId))
		_, err = memberFaker.Create(s)
		require.NoError(t, err)
		memberId = uint(memberFaker.Id.MemberId)
		return nil
	})
	require.NoError(t, err)

	return RepositoryOption{
		Name:       "PostgreSQL",
		Repository: NewCompetenceRepository(),
		Pool:       pool,
		TenantId:   tenantId,
		MemberId:   memberId,
	}
}
