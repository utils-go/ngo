package environment

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Environment provides information about, and means to manipulate, the current environment and platform
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.environment?view=netframework-4.7.2

// SpecialFolder identifies system special folders
type SpecialFolder int

const (
	// Desktop represents the logical Desktop rather than the physical file system location
	Desktop SpecialFolder = iota
	// Programs represents the directory that contains the user's program groups
	Programs
	// MyDocuments represents the My Documents folder
	MyDocuments
	// Personal represents the directory that serves as a common repository for documents
	Personal
	// Favorites represents the directory that serves as a common repository for the user's favorite items
	Favorites
	// Startup represents the directory that corresponds to the user's Startup program group
	Startup
	// Recent represents the directory that contains the user's most recently used documents
	Recent
	// SendTo represents the directory that contains the Send To menu items
	SendTo
	// StartMenu represents the directory that contains the Start menu items
	StartMenu
	// MyMusic represents the My Music folder
	MyMusic
	// MyVideos represents the My Videos folder
	MyVideos
	// DesktopDirectory represents the directory used to physically store file objects on the desktop
	DesktopDirectory
	// MyComputer represents the My Computer folder
	MyComputer
	// NetworkShortcuts represents the directory that contains network shortcuts
	NetworkShortcuts
	// Fonts represents the directory where fonts are stored
	Fonts
	// Templates represents the directory that serves as a common repository for document templates
	Templates
	// CommonStartMenu represents the directory for components that are shared across applications
	CommonStartMenu
	// CommonPrograms represents the directory for components that are shared across applications
	CommonPrograms
	// CommonStartup represents the directory for components that are shared across applications
	CommonStartup
	// CommonDesktopDirectory represents the directory for components that are shared across applications
	CommonDesktopDirectory
	// ApplicationData represents the directory that serves as a common repository for application-specific data for the current roaming user
	ApplicationData
	// PrinterShortcuts represents the directory that contains printer shortcuts
	PrinterShortcuts
	// LocalApplicationData represents the directory that serves as a common repository for application-specific data that is used by the current, non-roaming user
	LocalApplicationData
	// InternetCache represents the directory that serves as a common repository for temporary Internet files
	InternetCache
	// Cookies represents the directory that serves as a common repository for Internet cookies
	Cookies
	// History represents the directory that serves as a common repository for Internet history items
	History
	// CommonApplicationData represents the directory that serves as a common repository for application-specific data that is used by all users
	CommonApplicationData
	// Windows represents the Windows directory or SYSROOT
	Windows
	// System represents the System directory
	System
	// ProgramFiles represents the program files directory
	ProgramFiles
	// MyPictures represents the My Pictures folder
	MyPictures
	// UserProfile represents the user's profile folder
	UserProfile
	// SystemX86 represents the System directory on 64-bit systems
	SystemX86
	// ProgramFilesX86 represents the program files directory on 64-bit systems
	ProgramFilesX86
	// CommonProgramFiles represents the directory for components that are shared across applications
	CommonProgramFiles
	// CommonProgramFilesX86 represents the directory for components that are shared across applications on 64-bit systems
	CommonProgramFilesX86
	// CommonTemplates represents the directory for components that are shared across applications
	CommonTemplates
	// CommonDocuments represents the directory for components that are shared across applications
	CommonDocuments
	// CommonAdminTools represents the directory for components that are shared across applications
	CommonAdminTools
	// AdminTools represents the directory that serves as a common repository for administrative tools
	AdminTools
	// CommonMusic represents the directory for components that are shared across applications
	CommonMusic
	// CommonPictures represents the directory for components that are shared across applications
	CommonPictures
	// CommonVideos represents the directory for components that are shared across applications
	CommonVideos
	// Resources represents the file resources that are accessible to all users
	Resources
	// LocalizedResources represents the file resources that are accessible to all users
	LocalizedResources
	// CommonOemLinks represents the directory for components that are shared across applications
	CommonOemLinks
	// CDBurning represents the directory that serves as a staging area for files waiting to be written to CD
	CDBurning
)

// NewLine gets the newline string defined for this environment
func NewLine() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

// MachineName gets the NetBIOS name of this local computer
func MachineName() string {
	if hostname, err := os.Hostname(); err == nil {
		return hostname
	}
	return "localhost"
}

// OSVersion gets a string that identifies the operating system
func OSVersion() string {
	return runtime.GOOS + " " + runtime.GOARCH
}

// ProcessorCount gets the number of processors on the current machine
func ProcessorCount() int {
	return runtime.NumCPU()
}

// UserDomainName gets the network domain name associated with the current user
func UserDomainName() string {
	if runtime.GOOS == "windows" {
		if domain := os.Getenv("USERDOMAIN"); domain != "" {
			return domain
		}
	}
	return MachineName()
}

// UserName gets the user name of the person who is currently logged on to the operating system
func UserName() string {
	if currentUser, err := user.Current(); err == nil {
		return currentUser.Username
	}
	if username := os.Getenv("USER"); username != "" {
		return username
	}
	if username := os.Getenv("USERNAME"); username != "" {
		return username
	}
	return "unknown"
}

// Version gets the version of the common language runtime
func Version() string {
	return runtime.Version()
}

// WorkingSet gets the amount of physical memory mapped to the process context
func WorkingSet() int64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return int64(m.Sys)
}

// CurrentDirectory gets or sets the fully qualified path of the current working directory
func CurrentDirectory() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	return ""
}

// SetCurrentDirectory sets the current working directory
func SetCurrentDirectory(path string) error {
	return os.Chdir(path)
}

// SystemDirectory gets the fully qualified path of the system directory
func SystemDirectory() string {
	switch runtime.GOOS {
	case "windows":
		if systemRoot := os.Getenv("SYSTEMROOT"); systemRoot != "" {
			return filepath.Join(systemRoot, "System32")
		}
		return "C:\\Windows\\System32"
	case "linux", "darwin":
		return "/usr/bin"
	default:
		return "/bin"
	}
}

// GetEnvironmentVariable retrieves the value of an environment variable from the current process
func GetEnvironmentVariable(variable string) string {
	return os.Getenv(variable)
}

// GetEnvironmentVariables retrieves all environment variable names and their values from the current process
func GetEnvironmentVariables() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		if pair := strings.SplitN(e, "=", 2); len(pair) == 2 {
			env[pair[0]] = pair[1]
		}
	}
	return env
}

// SetEnvironmentVariable creates, modifies, or deletes an environment variable stored in the current process
func SetEnvironmentVariable(variable, value string) error {
	return os.Setenv(variable, value)
}

// ExpandEnvironmentVariables replaces the name of each environment variable embedded in the specified string with the string equivalent of the value of the variable
func ExpandEnvironmentVariables(name string) string {
	return os.ExpandEnv(name)
}

// GetCommandLineArgs returns a string array containing the command-line arguments for the current process
func GetCommandLineArgs() []string {
	return os.Args
}

// GetFolderPath gets the path to the system special folder that is identified by the specified enumeration
func GetFolderPath(folder SpecialFolder) string {
	currentUser, err := user.Current()
	if err != nil {
		return ""
	}

	switch folder {
	case Desktop, DesktopDirectory:
		return filepath.Join(currentUser.HomeDir, "Desktop")
	case MyDocuments, Personal:
		return filepath.Join(currentUser.HomeDir, "Documents")
	case MyMusic:
		return filepath.Join(currentUser.HomeDir, "Music")
	case MyPictures:
		return filepath.Join(currentUser.HomeDir, "Pictures")
	case MyVideos:
		return filepath.Join(currentUser.HomeDir, "Videos")
	case UserProfile:
		return currentUser.HomeDir
	case ApplicationData:
		if runtime.GOOS == "windows" {
			return os.Getenv("APPDATA")
		}
		return filepath.Join(currentUser.HomeDir, ".config")
	case LocalApplicationData:
		if runtime.GOOS == "windows" {
			return os.Getenv("LOCALAPPDATA")
		}
		return filepath.Join(currentUser.HomeDir, ".local", "share")
	case CommonApplicationData:
		if runtime.GOOS == "windows" {
			return os.Getenv("PROGRAMDATA")
		}
		return "/usr/share"
	case ProgramFiles:
		if runtime.GOOS == "windows" {
			if pf := os.Getenv("PROGRAMFILES"); pf != "" {
				return pf
			}
			return "C:\\Program Files"
		}
		return "/usr/bin"
	case ProgramFilesX86:
		if runtime.GOOS == "windows" {
			if pf := os.Getenv("PROGRAMFILES(X86)"); pf != "" {
				return pf
			}
			return "C:\\Program Files (x86)"
		}
		return "/usr/bin"
	case System:
		return SystemDirectory()
	case Windows:
		if runtime.GOOS == "windows" {
			if systemRoot := os.Getenv("SYSTEMROOT"); systemRoot != "" {
				return systemRoot
			}
			return "C:\\Windows"
		}
		return "/usr"
	case Fonts:
		if runtime.GOOS == "windows" {
			return filepath.Join(os.Getenv("SYSTEMROOT"), "Fonts")
		} else if runtime.GOOS == "darwin" {
			return "/System/Library/Fonts"
		}
		return "/usr/share/fonts"
	case Templates:
		return filepath.Join(currentUser.HomeDir, "Templates")
	case Startup:
		if runtime.GOOS == "windows" {
			return filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
		}
		return filepath.Join(currentUser.HomeDir, ".config", "autostart")
	default:
		return ""
	}
}

// Exit terminates this process and returns an exit code to the operating system
func Exit(exitCode int) {
	os.Exit(exitCode)
}

// ExitCode gets or sets the exit code of the process
var ExitCode int = 0

// HasShutdownStarted gets a value that indicates whether the current application domain is being unloaded or the common language runtime (CLR) is shutting down
func HasShutdownStarted() bool {
	// In Go, we don't have a direct equivalent, so we return false
	return false
}

// Is64BitOperatingSystem gets a value that indicates whether the current operating system is a 64-bit operating system
func Is64BitOperatingSystem() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
}

// Is64BitProcess gets a value that indicates whether the current process is a 64-bit process
func Is64BitProcess() bool {
	return strconv.IntSize == 64
}

// TickCount gets the number of milliseconds elapsed since the system started
func TickCount() int64 {
	// This is an approximation since Go doesn't have direct access to system uptime
	// In a real implementation, you might want to use platform-specific code
	return 0
}

// GetLogicalDrives returns an array of string containing the names of the logical drives on the current computer
func GetLogicalDrives() []string {
	var drives []string
	
	if runtime.GOOS == "windows" {
		// On Windows, check common drive letters
		for c := 'A'; c <= 'Z'; c++ {
			drive := string(c) + ":\\"
			if _, err := os.Stat(drive); err == nil {
				drives = append(drives, drive)
			}
		}
	} else {
		// On Unix-like systems, return common mount points
		drives = []string{"/"}
		// You could extend this to read /proc/mounts on Linux
	}
	
	return drives
}