package endorser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emacsway/grade/grade/internal/domain/endorser"
	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/domain/tenant"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	memberRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/member"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/seedwork/testutils"
	tenantRepo "github.com/emacsway/grade/grade/internal/infrastructure/repositories/tenant"
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
	var exporterActual endorser.EndorserExporter
	var exporterRead endorser.EndorserExporter
	factory := endorser.NewEndorserFaker(
		endorser.WithTenantId(repositoryOption.TenantId),
		endorser.WithMemberId(repositoryOption.MemberId),
	)
	agg, err := factory.Create()
	require.NoError(t, err)
	err = repositoryOption.Repository.Insert(agg)
	require.NoError(t, err)
	agg.Export(&exporterActual)

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
	Repository *EndorserRepository
	Session    infrastructure.DbSession
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
	session := infrastructure.NewPgxSession(db)
	tenantFaker := tenantRepo.NewTenantFaker(session)
	aTenant, err := tenantFaker.Create()
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
		Repository: NewEndorserRepository(session),
		Session:    session,
		TenantId:   uint(tenantExp.Id),
		MemberId:   uint(memberFaker.Id.MemberId),
	}
}
