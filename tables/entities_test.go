package tables

import (
  "testing"
)

func TestInvalidDatasetId(t *testing.T) {
  _, err := NewTableEntity("BadDataset", "testproject", "test")
  if err == nil {
    t.Error("Should have failed")
  }
}

func TestViewerSchema(t *testing.T) {
  table, err := NewTableEntity("Viewer_test", "testproject", "test")
  if err != nil {
    t.Error("Could not get a table entity for Viewer schema")
  }

  if table.TableReference.ProjectId != "testproject" {
    t.Error("Project id is incorrect")
  }

  if table.TableReference.TableId != "test" {
    t.Error("Table name is invalid")
  }

  if len(table.Schema.Fields) != 4 {
    t.Error("Viewer table contains incorrect field count")
  }
}

func TestCAPSchema(t *testing.T) {
  table, err := NewTableEntity("CAP_test", "testproject", "test")
  if err != nil {
    t.Error("Could not get a table entity for CAP schema")
  }

  if len(table.Schema.Fields) != 8 {
    t.Error("CAP table contains incorrect field count")
  }
}

func TestOLPSchema(t *testing.T) {
  table, err := NewTableEntity("OLP_test", "testproject", "test")
  if err != nil {
    t.Error("Could not get a table entity for OLP schema")
  }

  if len(table.Schema.Fields) != 7 {
    t.Error("OLP table contains incorrect field count")
  }
}
