package enum

// Status : TODO の状態を管理する
type Status int

const (
	_ Status = iota
	// NotExecuted : 未実行
	NotExecuted
	// Executing : 実行中
	Executing
	// Done : 完了
	Done
)

// fmt.Stringer インターフェース実装
func (s Status) String() string {
	switch s {
	case NotExecuted:
		return "未実行"
	case Executing:
		return "実行中"
	case Done:
		return "完了"
	default:
		return "Unknown"
	}
}
