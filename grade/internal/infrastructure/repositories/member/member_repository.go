package member

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/infrastructure"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/member/queries"
)

func NewMemberRepository(session infrastructure.DbSession) *MemberRepository {
	return &MemberRepository{
		session: session,
	}
}

type MemberRepository struct {
	session infrastructure.DbSession
}

func (r *MemberRepository) Insert(agg *member.Member) error {
	q := &queries.MemberInsertQuery{}
	agg.Export(q)
	result, err := q.Evaluate(r.session)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	if lastInsertId == 0 {
		return fmt.Errorf("wrong LastInsertId: %d", lastInsertId)
	}
	return nil
}

func (r *MemberRepository) Get(id memberVal.MemberId) (*member.Member, error) {
	q := queries.MemberGetQuery{Id: id}
	return q.Get(r.session)
}
