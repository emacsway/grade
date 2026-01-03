package values

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberIdEqual(t *testing.T) {
	t1 := uint(10)
	m1 := uint(3)
	t2 := uint(11)
	m2 := uint(4)
	cases := []struct {
		TenantId       uint
		MemberId       uint
		OtherTenantId  uint
		OtherMemberId  uint
		ExpectedResult bool
	}{
		{t1, m2, t1, m2, true},
		{t1, m1, t1, m2, false},
		{t2, m2, t1, m2, false},
		{t1, m2, t1, m1, false},
		{t1, m2, t2, m2, false},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			id, err := NewMemberId(c.TenantId, c.MemberId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			otherId, err := NewMemberId(c.OtherTenantId, c.OtherMemberId)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			r := id.Equal(otherId)
			assert.Equal(t, c.ExpectedResult, r)
		})
	}
}

func TestMemberIdExport(t *testing.T) {
	var actualExporter MemberIdExporter
	f := NewMemberIdFaker()
	cid, err := f.Create()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	cid.Export(&actualExporter)
	assert.Equal(t, MemberIdExporter{
		TenantId: f.TenantId,
		MemberId: f.MemberId,
	}, actualExporter)
}
