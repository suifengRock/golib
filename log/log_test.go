package log

import (
	"testing"
)

var tmp string = "SELECT `id`, `father_id`, `tpl_id`, `project_id`, `name`, `desc`, `type`, `cls`, `path`, `create_at`, `update_at`, `delete_at` FROM `category` WHERE (name = ?) AND (father_id = ?) AND (`delete_at` IS NULL or `delete_at` = '0001-01-01 00:00:00') ORDER BY `update_at` DESC LIMIT 1 [args] [月度检查 0];" +
	"SELECT `id`, `father_id`, `tpl_id`, `project_id`, `name`, `desc`, `type`, `cls`, `path`, `create_at`, `update_at`, `delete_at` FROM `category` WHERE (name = ?) AND (father_id = ?) AND (`delete_at` IS NULL or `delete_at` = '0001-01-01 00:00:00') ORDER BY `update_at` DESC LIMIT 1 [args] [月度检查 0];" +
	"SELECT `id`, `father_id`, `tpl_id`, `project_id`, `name`, `desc`, `type`, `cls`, `path`, `create_at`, `update_at`, `delete_at` FROM `category` WHERE (name = ?) AND (father_id = ?) AND (`delete_at` IS NULL or `delete_at` = '0001-01-01 00:00:00') ORDER BY `update_at` DESC LIMIT 1 [args] [月度检查 0];"

func TestInfo(t *testing.T) {
	logger.Info(tmp)
}

func TestDebug(t *testing.T) {
	logger.Debug(tmp)
}

func testWarning(t *testing.T) {
	logger.Warning(tmp)
}

func BenchmarkInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("%s..........%d", tmp, i)
	}
}

func BenchmarkDebug(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("..........%d", i)
	}
}
