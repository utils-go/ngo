package environment

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestMachineName(t *testing.T) {
	name := MachineName()
	if name == "" {
		t.Error("Expected non-empty machine name")
	}
}

func TestOSVersion(t *testing.T) {
	version := OSVersion()
	if version == "" {
		t.Error("Expected non-empty OS version")
	}
	
	// Should contain OS and architecture
	if !contains(version, runtime.GOOS) {
		t.Errorf("Expected OS version to contain '%s', got '%s'", runtime.GOOS, version)
	}
}

func TestProcessorCount(t *testing.T) {
	count := ProcessorCount()
	if count <= 0 {
		t.Errorf("Expected positive processor count, got %d", count)
	}
}

func TestUserName(t *testing.T) {
	name := UserName()
	if name == "" {
		t.Error("Expected non-empty user name")
	}
}

func TestCurrentDirectory(t *testing.T) {
	dir := CurrentDirectory()
	if dir == "" {
		t.Error("Expected non-empty current directory")
	}
}

func TestGetEnvironmentVariable(t *testing.T) {
	// Set a test environment variable
	testKey := "NGO_TEST_VAR"
	testValue := "test_value"
	
	err := SetEnvironmentVariable(testKey, testValue)
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	
	// Get the variable
	value := GetEnvironmentVariable(testKey)
	if value != testValue {
		t.Errorf("Expected '%s', got '%s'", testValue, value)
	}
	
	// Clean up
	os.Unsetenv(testKey)
}

func TestGetEnvironmentVariables(t *testing.T) {
	vars := GetEnvironmentVariables()
	if len(vars) == 0 {
		t.Error("Expected some environment variables")
	}
	
	// Should contain PATH (or Path on Windows)
	hasPath := false
	for key := range vars {
		if key == "PATH" || key == "Path" {
			hasPath = true
			break
		}
	}
	if !hasPath {
		t.Error("Expected environment variables to contain PATH")
	}
}

func TestExpandEnvironmentVariables(t *testing.T) {
	// Set a test variable
	testKey := "NGO_TEST_EXPAND"
	testValue := "expanded_value"
	
	SetEnvironmentVariable(testKey, testValue)
	defer os.Unsetenv(testKey)
	
	// Test expansion
	input := "prefix_${NGO_TEST_EXPAND}_suffix"
	result := ExpandEnvironmentVariables(input)
	expected := "prefix_expanded_value_suffix"
	
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestGetCommandLineArgs(t *testing.T) {
	args := GetCommandLineArgs()
	if len(args) == 0 {
		t.Error("Expected at least one command line argument (program name)")
	}
}

func TestGetFolderPath(t *testing.T) {
	// Test user profile folder
	userProfile := GetFolderPath(UserProfile)
	if userProfile == "" {
		t.Error("Expected non-empty user profile path")
	}
	
	// Test desktop folder
	desktop := GetFolderPath(Desktop)
	if desktop == "" {
		t.Error("Expected non-empty desktop path")
	}
}

func TestIs64BitProcess(t *testing.T) {
	is64Bit := Is64BitProcess()
	// This should match the architecture we're running on
	expected := runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
	if is64Bit != expected {
		t.Errorf("Expected Is64BitProcess() to be %t, got %t", expected, is64Bit)
	}
}

func TestIs64BitOperatingSystem(t *testing.T) {
	is64BitOS := Is64BitOperatingSystem()
	// This should match the architecture
	expected := runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
	if is64BitOS != expected {
		t.Errorf("Expected Is64BitOperatingSystem() to be %t, got %t", expected, is64BitOS)
	}
}

func TestNewLine(t *testing.T) {
	newline := NewLine()
	if runtime.GOOS == "windows" {
		if newline != "\r\n" {
			t.Errorf("Expected Windows newline '\\r\\n', got '%s'", newline)
		}
	} else {
		if newline != "\n" {
			t.Errorf("Expected Unix newline '\\n', got '%s'", newline)
		}
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && 
			(s[:len(substr)] == substr || 
			 s[len(s)-len(substr):] == substr || 
			 strings.Contains(s, substr))))
}