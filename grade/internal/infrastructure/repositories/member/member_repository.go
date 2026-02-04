package member

import (
	"fmt"

	"github.com/emacsway/grade/grade/internal/domain/member"
	memberVal "github.com/emacsway/grade/grade/internal/domain/member/values"
	"github.com/emacsway/grade/grade/internal/infrastructure/repositories/member/queries"
	"github.com/krew-solutions/ascetic-ddd-go/asceticddd/session"
)

func NewMemberRepository() *MemberRepository {
	return &MemberRepository{}
}

type MemberRepository struct{}

func (r *MemberRepository) Insert(s session.Session, agg *member.Member) error {
	q := &queries.MemberInsertQuery{}
	agg.Export(q)
	result, err := q.Evaluate(s)
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

func (r *MemberRepository) Get(s session.Session, id memberVal.MemberId) (*member.Member, error) {
	q := queries.MemberGetQuery{Id: id}
	return q.Get(s)
}
