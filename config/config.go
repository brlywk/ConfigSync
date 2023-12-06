package config

type (
	// Represents content of configuration file
	Config struct {
		SyncDirectory      string
		IncludeSubfolders  bool
		ExcludeHasPriority bool
		MappingFile        string
		CreateDirectories  bool
		SkipOnError        bool
		NeverDelete        bool
		Include            `toml:"include"`
		Exclude            `toml:"exclude"`
		Git                `toml:"git"`
	}

	// Files, directories and patterns to include
	Include struct {
		Files       []string
		Directories []string
		Patterns    []string
	}

	// Files, directories and patterns to exclude
	Exclude struct {
		Files       []string
		Directories []string
		Patterns    []string
	}

	// Git settings
	Git struct {
		Repo             string
		AlwaysPushOnSync bool
		CreateBranch     bool
	}
)
