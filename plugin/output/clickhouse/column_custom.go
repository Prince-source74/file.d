package clickhouse

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/ClickHouse/ch-go/proto"
	insaneJSON "github.com/vitkovskii/insane-json"
)

// ColDateTime represents Clickhouse DateTime type.
type ColDateTime struct {
	col *proto.ColDateTime
}

func NewColDateTime(col *proto.ColDateTime) *ColDateTime {
	return &ColDateTime{
		col: col,
	}
}

func (t *ColDateTime) Append(node *insaneJSON.StrictNode) error {
	if node == nil || node.IsNull() {
		return ErrNodeIsNil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}

	val := time.Unix(int64(v), 0)

	t.col.Append(val)

	return nil
}

// ColDateTime64 represents Clickhouse DateTime64 type.
type ColDateTime64 struct {
	col   *proto.ColDateTime64
	scale int64
}

func NewColDateTime64(col *proto.ColDateTime64, scale int64) *ColDateTime64 {
	return &ColDateTime64{
		col:   col,
		scale: scale,
	}
}

func (t *ColDateTime64) Append(node *insaneJSON.StrictNode) error {
	if node == nil || node.IsNull() {
		return ErrNodeIsNil
	}

	v, err := node.AsInt64()
	if err != nil {
		return err
	}

	nsec := v * t.scale
	val := time.Unix(nsec/1e9, nsec%1e9)

	t.col.Append(val)

	return nil
}

// ColIPv4 represents Clickhouse IPv4 type.
type ColIPv4 struct {
	col      *proto.ColIPv4
	nullCol  *proto.ColNullable[proto.IPv4]
	nullable bool
}

func NewColIPv4(nullable bool) *ColIPv4 {
	return &ColIPv4{
		col:      new(proto.ColIPv4),
		nullCol:  new(proto.ColIPv4).Nullable(),
		nullable: nullable,
	}
}

func (t *ColIPv4) Append(node *insaneJSON.StrictNode) error {
	if node == nil || node.IsNull() {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[proto.IPv4]())
		return nil
	}

	v, err := node.AsString()
	if err != nil {
		return err
	}

	addr, err := netip.ParseAddr(v)
	if err != nil {
		return err
	}
	if !addr.Is4() {
		return fmt.Errorf("invalid IPv6 value, val=%s", v)
	}

	val := proto.ToIPv4(addr)

	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

// ColIPv6 represents Clickhouse IPv6 type.
type ColIPv6 struct {
	col      *proto.ColIPv6
	nullCol  *proto.ColNullable[proto.IPv6]
	nullable bool
}

func NewColIPv6(nullable bool) *ColIPv6 {
	return &ColIPv6{
		col:      new(proto.ColIPv6),
		nullCol:  new(proto.ColIPv6).Nullable(),
		nullable: nullable,
	}
}

func (t *ColIPv6) Append(node *insaneJSON.StrictNode) error {
	if node == nil || node.IsNull() {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[proto.IPv6]())
		return nil
	}

	v, err := node.AsString()
	if err != nil {
		return err
	}

	addr, err := netip.ParseAddr(v)
	if err != nil {
		return err
	}
	if !addr.Is6() {
		return fmt.Errorf("invalid IPv6 value, val=%s", v)
	}
	val := proto.ToIPv6(addr)

	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

// ColEnum8 represents Clickhouse Enum8 type.
type ColEnum8 struct {
	col *proto.ColEnum
}

func NewColEnum8(col *proto.ColEnum) *ColEnum8 {
	return &ColEnum8{
		col: col,
	}
}

func (t *ColEnum8) Append(node *insaneJSON.StrictNode) error {
	if node == nil || node.IsNull() {
		return ErrNodeIsNil
	}
	val, err := node.AsString()
	if err != nil {
		return err
	}
	t.col.Append(val)

	return nil
}

// ColEnum16 represents Clickhouse Enum16 type.
type ColEnum16 struct {
	col *proto.ColEnum
}

func NewColEnum16(col *proto.ColEnum) *ColEnum8 {
	return &ColEnum8{
		col: col,
	}
}

func (t *ColEnum16) Append(node *insaneJSON.StrictNode) error {
	if node == nil || node.IsNull() {
		return ErrNodeIsNil
	}
	val, err := node.AsString()
	if err != nil {
		return err
	}
	t.col.Append(val)

	return nil
}
