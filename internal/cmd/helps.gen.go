// Code generated by github.com/twpayne/chezmoi/internal/cmds/generate-helps. DO NOT EDIT.

package cmd

import (
	"github.com/twpayne/chezmoi/v2/internal/chezmoiset"
)

type help struct {
	longHelp   string
	example    string
	longFlags  chezmoiset.Set[string]
	shortFlags chezmoiset.Set[string]
}

var helps = map[string]*help{
	"add": {
		longHelp: "" +
			"Description:\n" +
			"  Add targets to the source state. If any target is already in the source state,\n" +
			"  then its source state is replaced with its current state in the destination\n" +
			"  directory.",
		example: "" +
			"  chezmoi add ~/.bashrc\n" +
			"  chezmoi add ~/.gitconfig --template\n" +
			"  chezmoi add ~/.ssh/id_rsa --encrypt\n" +
			"  chezmoi add ~/.vim --recursive\n" +
			"  chezmoi add ~/.oh-my-zsh --exact --recursive",
		longFlags: chezmoiset.New(
			"autotemplate",
			"create",
			"encrypt",
			"exact",
			"exclude",
			"follow",
			"force",
			"include",
			"prompt",
			"quiet",
			"recursive",
			"secrets",
			"template",
			"template-symlinks",
		),
		shortFlags: chezmoiset.New(
			"T",
			"a",
			"f",
			"i",
			"p",
			"q",
			"r",
			"x",
		),
	},
	"age": {
		longHelp: "" +
			"Description:\n" +
			"  Interact with age's passphrase-based encryption.",
		example: "" +
			"  chezmoi age encrypt --passphrase plaintext.txt > ciphertext.txt\n" +
			"  chezmoi age decrypt --passphrase ciphertext.txt > decrypted-ciphertext.txt",
	},
	"apply": {
		longHelp: "" +
			"Description:\n" +
			"  Ensure that target... are in the target state, updating them if necessary. If no\n" +
			"  targets are specified, the state of all targets are ensured. If a target has\n" +
			"  been modified since chezmoi last wrote it then the user will be prompted if they\n" +
			"  want to overwrite the file.",
		example: "" +
			"  chezmoi apply\n" +
			"  chezmoi apply --dry-run --verbose\n" +
			"  chezmoi apply ~/.bashrc",
		longFlags: chezmoiset.New(
			"exclude",
			"include",
			"init",
			"parent-dirs",
			"recursive",
			"source-path",
		),
		shortFlags: chezmoiset.New(
			"P",
			"i",
			"r",
			"x",
		),
	},
	"archive": {
		longHelp: "" +
			"Description:\n" +
			"  Generate an archive of the target state, or only the targets specified. This can\n" +
			"  be piped into tar to inspect the target state.",
		example: "" +
			"  chezmoi archive | tar tvf -\n" +
			"  chezmoi archive --output=dotfiles.tar.gz\n" +
			"  chezmoi archive --output=dotfiles.zip",
		longFlags: chezmoiset.New(
			"exclude",
			"format",
			"gzip",
			"include",
			"init",
			"parent-dirs",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"P",
			"f",
			"i",
			"r",
			"x",
			"z",
		),
	},
	"cat": {
		longHelp: "" +
			"Description:\n" +
			"  Write the target contents of targets to stdout. targets must be files, scripts,\n" +
			"  or symlinks. For files, the target file contents are written. For scripts, the\n" +
			"  script's contents are written. For symlinks, the target is written.",
		example: "" +
			"  chezmoi cat ~/.bashrc",
	},
	"cat-config": {
		longHelp: "" +
			"Description:\n" +
			"  Print the configuration file.",
		example: "" +
			"  chezmoi cat-config",
	},
	"cd": {
		longHelp: "" +
			"Description:\n" +
			"  Launch a shell in the working tree (typically the source directory). chezmoi\n" +
			"  will launch the command set by the cd.command configuration variable with any\n" +
			"  extra arguments specified by cd.args. If this is not set, chezmoi will attempt\n" +
			"  to detect your shell and finally fall back to an OS-specific default.\n" +
			"\n" +
			"  If the optional argument path is present, the shell will be launched in the\n" +
			"  source directory corresponding to path.\n" +
			"\n" +
			"  The shell will have various CHEZMOI* environment variables set, as for scripts.",
		example: "" +
			"  chezmoi cd\n" +
			"  chezmoi cd ~\n" +
			"  chezmoi cd ~/.config",
	},
	"chattr": {
		longHelp: "" +
			"Description:\n" +
			"  Change the attributes and/or type of targets. modifier specifies what to modify.\n" +
			"\n" +
			"  Add attributes by specifying them or their abbreviations directly, optionally\n" +
			"  prefixed with a plus sign (+). Remove attributes by prefixing them or their\n" +
			"  attributes with the string no or a minus sign (-). The available attribute\n" +
			"  modifiers and their abbreviations are:\n" +
			"\n" +
			"  Attribute modifier                         │Abbreviation\n" +
			"  ───────────────────────────────────────────┼────────────────────────────────────\n" +
			"  after                                      │a\n" +
			"  before                                     │b\n" +
			"  empty                                      │e\n" +
			"  encrypted                                  │none\n" +
			"  exact                                      │none\n" +
			"  executable                                 │x\n" +
			"  external                                   │none\n" +
			"  once                                       │o\n" +
			"  private                                    │p\n" +
			"  readonly                                   │r\n" +
			"  remove                                     │none\n" +
			"  template                                   │t\n" +
			"\n" +
			"  The type of a target can be changed using a type modifier:\n" +
			"\n" +
			"  Type modifier\n" +
			"  ────────────────────────────────────────────────────────────────────────────────\n" +
			"  create\n" +
			"  modify\n" +
			"  script\n" +
			"  symlink\n" +
			"\n" +
			"  The negative form of type modifiers, e.g. nocreate, changes the target to be a\n" +
			"  regular file if it is of that type, otherwise the type is left unchanged.\n" +
			"\n" +
			"  Multiple modifications may be specified by separating them with a comma (,). If\n" +
			"  you use the -modifier form then you must put modifier after a -- to prevent\n" +
			"  chezmoi\n" +
			"  from interpreting -modifier as an option.",
		example: "" +
			"  chezmoi chattr template ~/.bashrc\n" +
			"  chezmoi chattr noempty ~/.profile\n" +
			"  chezmoi chattr private,template ~/.netrc\n" +
			"  chezmoi chattr -- -x ~/.zshrc\n" +
			"  chezmoi chattr +create,+private ~/.kube/config",
		longFlags: chezmoiset.New(
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"r",
		),
	},
	"completion": {
		longHelp: "" +
			"Description:\n" +
			"  Generate shell completion code for the specified shell (bash, fish, powershell,\n" +
			"  or zsh).",
		example: "" +
			"  chezmoi completion bash\n" +
			"  chezmoi completion fish --output=~/.config/fish/completions/chezmoi.fish",
	},
	"data": {
		longHelp: "" +
			"Description:\n" +
			"  Write the computed template data to stdout.",
		example: "" +
			"  chezmoi data\n" +
			"  chezmoi data --format=yaml",
		longFlags: chezmoiset.New(
			"format",
		),
		shortFlags: chezmoiset.New(
			"f",
		),
	},
	"decrypt": {
		longHelp: "" +
			"Description:\n" +
			"  Decrypt files using chezmoi's configured encryption. If no files are given,\n" +
			"  decrypt the standard input. The decrypted result is written to the standard\n" +
			"  output or a file if the --output flag is set.",
	},
	"destroy": {
		longHelp: "" +
			"Description:\n" +
			"",
		longFlags: chezmoiset.New(
			"force",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"r",
		),
	},
	"diff": {
		longHelp: "" +
			"Description:\n" +
			"  Print the difference between the target state and the destination state for\n" +
			"  targets. If no targets are specified, print the differences for all targets.\n" +
			"\n" +
			"  If a diff.pager command is set in the configuration file then the output will be\n" +
			"  piped into it.\n" +
			"\n" +
			"  If diff.command is set then it will be invoked to show individual file\n" +
			"  differences with diff.args passed as arguments. Each element of diff.args is\n" +
			"  interpreted as a template with the variables .Destination and .Target available\n" +
			"  corresponding to the path of the file in the source and target state\n" +
			"  respectively. The default value of diff.args is [\"{{ .Destination }}\", \"{{\n" +
			"  .Target }}\"]. If diff.args does not contain any template arguments then {{\n" +
			"  .Destination }} and {{ .Target }} will be appended automatically.",
		example: "" +
			"  chezmoi diff\n" +
			"  chezmoi diff ~/.bashrc",
		longFlags: chezmoiset.New(
			"exclude",
			"include",
			"init",
			"pager",
			"parent-dirs",
			"recursive",
			"reverse",
			"script-contents",
		),
		shortFlags: chezmoiset.New(
			"P",
			"i",
			"r",
			"x",
		),
	},
	"doctor": {
		longHelp: "" +
			"Description:\n" +
			"  Check for potential problems.",
		example: "" +
			"  chezmoi doctor",
		longFlags: chezmoiset.New(
			"no-network",
		),
	},
	"dump": {
		longHelp: "" +
			"Description:\n" +
			"  Dump the target state of targets. If no targets are specified, then the entire\n" +
			"  target state.",
		example: "" +
			"  chezmoi dump ~/.bashrc\n" +
			"  chezmoi dump --format=yaml",
		longFlags: chezmoiset.New(
			"exclude",
			"format",
			"include",
			"init",
			"parent-dirs",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"P",
			"f",
			"i",
			"r",
			"x",
		),
	},
	"dump-config": {
		longHelp: "" +
			"Description:\n" +
			"  Dump the configuration.",
		example: "" +
			"  chezmoi dump-config",
		longFlags: chezmoiset.New(
			"format",
		),
		shortFlags: chezmoiset.New(
			"f",
		),
	},
	"edit": {
		longHelp: "" +
			"Description:\n" +
			"  Edit the source state of targets, which must be files or symlinks. If no targets\n" +
			"  are given then the working tree of the source directory is opened.\n" +
			"\n" +
			"  Encrypted files are decrypted to a private temporary directory and the editor is\n" +
			"  invoked with the decrypted file. When the editor exits the edited decrypted file\n" +
			"  is re-encrypted and replaces the original file in the source state.\n" +
			"\n" +
			"  If the operating system supports hard links, then the edit command invokes the\n" +
			"  editor with filenames which match the target filename, unless the edit.hardlink\n" +
			"  configuration variable is set to false or the --hardlink=false command line flag\n" +
			"  is set.",
		example: "" +
			"  chezmoi edit ~/.bashrc\n" +
			"  chezmoi edit ~/.bashrc --apply\n" +
			"  chezmoi edit",
		longFlags: chezmoiset.New(
			"apply",
			"exclude",
			"hardlink",
			"include",
			"init",
			"watch",
		),
		shortFlags: chezmoiset.New(
			"a",
			"i",
			"x",
		),
	},
	"edit-config": {
		longHelp: "" +
			"Description:\n" +
			"  Edit the configuration file.",
		example: "" +
			"  chezmoi edit-config",
	},
	"edit-config-template": {
		longHelp: "" +
			"Description:\n" +
			"  Edit the configuration file template. If no configuration file template exists,\n" +
			"  then a new one is created with the contents of the current config file.",
		example: "" +
			"  chezmoi edit-config-template",
	},
	"encrypt": {
		longHelp: "" +
			"Description:\n" +
			"  Encrypt files using chezmoi's configured encryption. If no files are given,\n" +
			"  encrypt the standard input. The encrypted result is written to the standard\n" +
			"  output or a file if the --output flag is set.",
	},
	"execute-template": {
		longHelp: "" +
			"Description:\n" +
			"  Execute templates. This is useful for testing templates or for calling chezmoi\n" +
			"  from other scripts. templates are interpreted as literal templates, with no\n" +
			"  whitespace added to the output between arguments. If no templates are specified,\n" +
			"  the template is read from stdin.",
		example: "" +
			"  chezmoi execute-template '{{ .chezmoi.sourceDir }}'\n" +
			"  chezmoi execute-template '{{ .chezmoi.os }}' / '{{ .chezmoi.arch }}'\n" +
			"  echo '{{ .chezmoi | toJson }}' | chezmoi execute-template\n" +
			"  chezmoi execute-template --init --promptString email=me@home.org < ~/.\n" +
			"local/share/chezmoi/.chezmoi.toml.tmpl",
		longFlags: chezmoiset.New(
			"init",
			"left-delimiter",
			"promptBool",
			"promptChoice",
			"promptInt",
			"promptString",
			"right-delimiter",
			"stdinisatty",
			"with-stdin",
		),
		shortFlags: chezmoiset.New(
			"i",
			"p",
		),
	},
	"forget": {
		longHelp: "" +
			"Description:\n" +
			"  Remove targets from the source state, i.e. stop managing them. targets must have\n" +
			"  entries in the source state. They cannot be externals.",
		example: "" +
			"  chezmoi forget ~/.bashrc",
	},
	"generate": {
		longHelp: "" +
			"Description:\n" +
			"  Generates output for use with chezmoi. The currently supported outputs are:\n" +
			"\n" +
			"  Output         │Description\n" +
			"  ───────────────┼────────────────────────────────────────────────────────────────\n" +
			"  git-commit-mes…│A git commit message, describing the changes to the source dire…\n" +
			"  install.sh     │An install script, suitable for use with GitHub Codespaces",
		example: "" +
			"  chezmoi generate install.sh > install.sh\n" +
			"  chezmoi git commit -m \"$(chezmoi generate git-commit-message)\"",
	},
	"git": {
		longHelp: "" +
			"Description:\n" +
			"  Run git args in the working tree (typically the source directory).",
		example: "" +
			"  chezmoi git add .\n" +
			"  chezmoi git add dot_gitconfig\n" +
			"  chezmoi git -- commit -m \"Add .gitconfig\"",
	},
	"help": {
		longHelp: "" +
			"Description:\n" +
			"  Print the help associated with command, or general help if no command is given.",
	},
	"ignored": {
		longHelp: "" +
			"Description:\n" +
			"  Print the list of entries ignored by chezmoi.",
		example: "" +
			"  chezmoi ignored",
		longFlags: chezmoiset.New(
			"nul-path-separator",
			"tree",
		),
		shortFlags: chezmoiset.New(
			"0",
			"t",
		),
	},
	"import": {
		longHelp: "" +
			"Description:\n" +
			"  Import the source state from an archive file in to a directory in the source\n" +
			"  state. This is primarily used to make subdirectories of your home directory\n" +
			"  exactly match the contents of a downloaded archive. You will generally always\n" +
			"  want to set the --destination, --exact, and --remove-destination flags.\n" +
			"\n" +
			"  The supported archive formats are tar, tar.gz, tgz, tar.bz2, tbz2, txz, tar.zst,\n" +
			"  and zip.",
		example: "" +
			"  curl -s -L -o ${TMPDIR}/oh-my-zsh-master.tar.gz https://github.\n" +
			"com/ohmyzsh/ohmyzsh/archive/master.tar.gz\n" +
			"  mkdir -p $(chezmoi source-path)/dot_oh-my-zsh\n" +
			"  chezmoi import --strip-components 1 --destination ~/.oh-my-zsh ${TMPDIR}/oh-my-\n" +
			"zsh-master.tar.gz",
		longFlags: chezmoiset.New(
			"destination",
			"exact",
			"exclude",
			"include",
			"remove-destination",
			"strip-components",
		),
		shortFlags: chezmoiset.New(
			"d",
			"i",
			"r",
			"x",
		),
	},
	"init": {
		longHelp: "" +
			"Description:\n" +
			"  Setup the source directory, generate the config file, and optionally update the\n" +
			"  destination directory to match the target state. This is done in the following\n" +
			"  order:\n" +
			"\n" +
			"  1. The source directory is initialized. If chezmoi does not detect a Git\n" +
			"  repository in the source directory, chezmoi will clone the provided repo\n" +
			"  into the source directory. If no repo is provided, chezmoi will initialize\n" +
			"  a new Git repository.\n" +
			"  2. If the initialized source directory contains a .chezmoi.$FORMAT.tmpl file,\n" +
			"  a new configuration file will be created using that file as a template.\n" +
			"  3. If the --apply flag is provided, chezmoi apply is run.\n" +
			"  4. If the --purge flag is provided, chezmoi will remove the source, config,\n" +
			"  and cache directories.\n" +
			"  5. If the --purge-binary is passed, chezmoi will attempt to remove its own\n" +
			"  binary.\n" +
			"\n" +
			"  By default, if repo is given, chezmoi will guess the full git repo URL, using\n" +
			"  HTTPS by default, or SSH if the --ssh option is specified, according to the\n" +
			"  following patterns:\n" +
			"\n" +
			"  Pattern     │HTTPS Repo                           │SSH repo\n" +
			"  ────────────┼─────────────────────────────────────┼─────────────────────────────\n" +
			"  user        │https://user@github.com/user/dotfile…│git@github.com:user/dotfiles…\n" +
			"  user/repo   │https://user@github.com/user/repo.git│git@github.com:user/repo.git\n" +
			"  site/user/r…│https://user@site/user/repo.git      │git@/user/repo.git\n" +
			"  sr.ht/~user │https://user@git.sr.ht/~user/dotfiles│git@git.sr.ht:~user/dotfiles…\n" +
			"  sr.ht/~user…│https://user@git.sr.ht/~user/repo    │git@git.sr.ht:~user/repo.git\n" +
			"\n" +
			"  To disable git repo URL guessing, pass the --guess-repo-url=false option.",
		example: "" +
			"  chezmoi init user\n" +
			"  chezmoi init user --apply\n" +
			"  chezmoi init user --apply --purge\n" +
			"  chezmoi init user/dots\n" +
			"  chezmoi init codeberg.org/user\n" +
			"  chezmoi init gitlab.com/user",
		longFlags: chezmoiset.New(
			"apply",
			"branch",
			"config-path",
			"data",
			"depth",
			"exclude",
			"git-lfs",
			"guess-repo-url",
			"include",
			"one-shot",
			"prompt",
			"promptBool",
			"promptChoice",
			"promptDefaults",
			"promptInt",
			"promptString",
			"purge",
			"purge-binary",
			"recurse-submodules",
			"ssh",
		),
		shortFlags: chezmoiset.New(
			"C",
			"P",
			"a",
			"d",
			"g",
			"i",
			"p",
			"x",
		),
	},
	"license": {
		longHelp: "" +
			"Description:\n" +
			"  Print chezmoi's license.",
		example: "" +
			"  chezmoi license",
	},
	"list": {
		longHelp: "" +
			"Description:\n" +
			"  list is an alias for managed.",
	},
	"manage": {
		longHelp: "" +
			"Description:\n" +
			"  manage is an alias for add for symmetry with unmanage.",
	},
	"managed": {
		longHelp: "" +
			"Description:\n" +
			"  List all managed entries in the destination directory under all paths in\n" +
			"  alphabetical order. When no paths are supplied, list all managed entries in the\n" +
			"  destination directory in alphabetical order.",
		example: "" +
			"  chezmoi managed\n" +
			"  chezmoi managed --include=files\n" +
			"  chezmoi managed --include=files,symlinks\n" +
			"  chezmoi managed -i dirs\n" +
			"  chezmoi managed -i dirs,files\n" +
			"  chezmoi managed -i files ~/.config\n" +
			"  chezmoi managed --exclude=encrypted --path-style=source-relative",
		longFlags: chezmoiset.New(
			"exclude",
			"format",
			"include",
			"nul-path-separator",
			"path-style",
			"tree",
		),
		shortFlags: chezmoiset.New(
			"0",
			"f",
			"i",
			"p",
			"t",
			"x",
		),
	},
	"merge": {
		longHelp: "" +
			"Description:\n" +
			"  Perform a three-way merge between the destination state, the target state, and\n" +
			"  the source state for each target. The merge tool is defined by the merge.command\n" +
			"  configuration variable, and defaults to vimdiff. If multiple targets are\n" +
			"  specified the merge tool is invoked separately and sequentially for each target.\n" +
			"  If the target state cannot be computed (for example if source is a template\n" +
			"  containing errors or an encrypted file that cannot be decrypted) a two-way merge\n" +
			"  is performed instead.\n" +
			"\n" +
			"  The order of arguments to merge.command is set by merge.args. Each argument is\n" +
			"  interpreted as a template with the variables .Destination, .Source, and .Target\n" +
			"  available corresponding to the path of the file in the destination state, the\n" +
			"  source state, and the target state respectively. The default value of merge.args\n" +
			"  is [\"{{ .Destination }}\", \"{{ .Source }}\", \"{{ .Target }}\"]. If merge.args does\n" +
			"  not contain any template arguments then {{ .Destination }}, {{ .Source }}, and\n" +
			"  {{ .Target }} will be appended automatically.",
		example: "" +
			"  chezmoi merge ~/.bashrc",
	},
	"merge-all": {
		longHelp: "" +
			"Description:\n" +
			"  Perform a three-way merge for file whose actual state does not match its target\n" +
			"  state. The merge is performed with chezmoi merge.",
		example: "" +
			"  chezmoi merge-all",
		longFlags: chezmoiset.New(
			"init",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"r",
		),
	},
	"purge": {
		longHelp: "" +
			"Description:\n" +
			"  Remove chezmoi's configuration, state, and source directory, but leave the\n" +
			"  target state intact.",
		example: "" +
			"  chezmoi purge\n" +
			"  chezmoi purge --force",
		longFlags: chezmoiset.New(
			"binary",
			"force",
		),
		shortFlags: chezmoiset.New(
			"P",
		),
	},
	"re-add": {
		longHelp: "" +
			"Description:\n" +
			"  Re-add modified files in the target state, preserving any encrypted_ attributes.\n" +
			"  chezmoi will not overwrite templates, and all entries that are not files are\n" +
			"  ignored. Directories are recursed into by default.\n" +
			"\n" +
			"  If no targets are specified then all modified files are re-added. If one or more\n" +
			"  targets are given then only those targets are re-added.",
		example: "" +
			"  chezmoi re-add\n" +
			"  chezmoi re-add ~/.bashrc\n" +
			"  chezmoi re-add --recursive=false ~/.config/git",
		longFlags: chezmoiset.New(
			"exclude",
			"include",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"i",
			"r",
			"x",
		),
	},
	"remove": {
		longHelp: "" +
			"Description:\n" +
			"  The remove command has been removed. Use the forget command or the destroy\n" +
			"  command instead.",
	},
	"rm": {
		longHelp: "" +
			"Description:\n" +
			"  The rm command has been removed. Use the forget command or the destroy command\n" +
			"  instead.",
	},
	"secret": {
		longHelp: "" +
			"Description:\n" +
			"  Run a secret manager's CLI, passing any extra arguments to the secret manager's\n" +
			"  CLI. This is primarily for verifying chezmoi's integration with a custom secret\n" +
			"  manager. Normally you would use chezmoi's existing template functions to\n" +
			"  retrieve secrets.",
		example: "" +
			"  chezmoi secret keyring set --service=service --user=user --value=password\n" +
			"  chezmoi secret keyring get --service=service --user=user\n" +
			"  chezmoi secret keyring delete --service=service --user=user",
	},
	"source-path": {
		longHelp: "" +
			"Description:\n" +
			"  Print the path to each target's source state. If no targets are specified then\n" +
			"  print the source directory.",
		example: "" +
			"  chezmoi source-path\n" +
			"  chezmoi source-path ~/.bashrc",
	},
	"state": {
		longHelp: "" +
			"Description:\n" +
			"  Manipulate the persistent state.",
		example: "" +
			"  chezmoi state data\n" +
			"  chezmoi state delete --bucket=$BUCKET --key=$KEY\n" +
			"  chezmoi state delete-bucket --bucket=$BUCKET\n" +
			"  chezmoi state dump\n" +
			"  chezmoi state get --bucket=$BUCKET --key=$KEY\n" +
			"  chezmoi state get-bucket --bucket=$BUCKET\n" +
			"  chezmoi state set --bucket=$BUCKET --key=$KEY --value=$VALUE\n" +
			"  chezmoi state reset",
	},
	"status": {
		longHelp: "" +
			"Description:\n" +
			"  Print the status of the files and scripts managed by chezmoi in a format similar\n" +
			"  to git status.\n" +
			"\n" +
			"  The first column of output indicates the difference between the last state\n" +
			"  written by chezmoi and the actual state. The second column indicates the\n" +
			"  difference between the actual state and the target state, and what effect\n" +
			"  running chezmoi apply will have.\n" +
			"\n" +
			"  Character     │Meaning       │First column           │Second column\n" +
			"  ──────────────┼──────────────┼───────────────────────┼──────────────────────────\n" +
			"  Space         │No change     │No change              │No change\n" +
			"  A             │Added         │Entry was created      │Entry will be created\n" +
			"  D             │Deleted       │Entry was deleted      │Entry will be deleted\n" +
			"  M             │Modified      │Entry was modified     │Entry will be modified\n" +
			"  R             │Run           │Not applicable         │Script will be run",
		example: "" +
			"  chezmoi status",
		longFlags: chezmoiset.New(
			"exclude",
			"include",
			"init",
			"parent-dirs",
			"path-style",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"P",
			"i",
			"p",
			"r",
			"x",
		),
	},
	"target-path": {
		longHelp: "" +
			"Description:\n" +
			"  Print the target path of each source path. If no source paths are specified then\n" +
			"  print the target directory.",
		example: "" +
			"  chezmoi target-path\n" +
			"  chezmoi target-path ~/.local/share/chezmoi/dot_zshrc",
	},
	"unmanage": {
		longHelp: "" +
			"Description:\n" +
			"  unmanage is an alias for forget for symmetry with manage.",
	},
	"unmanaged": {
		longHelp: "" +
			"Description:\n" +
			"  List all unmanaged files in paths. When no paths are supplied, list all\n" +
			"  unmanaged files in the destination directory.\n" +
			"\n" +
			"  It is an error to supply paths that are not found on the file system.",
		example: "" +
			"  chezmoi unmanaged\n" +
			"  chezmoi unmanaged ~/.config/chezmoi ~/.ssh",
		longFlags: chezmoiset.New(
			"nul-path-separator",
			"path-style",
			"tree",
		),
		shortFlags: chezmoiset.New(
			"0",
			"p",
			"t",
		),
	},
	"update": {
		longHelp: "" +
			"Description:\n" +
			"  Pull changes from the source repo and apply any changes.\n" +
			"\n" +
			"  If update.command is set then chezmoi will run update.command with update.args\n" +
			"  in the working tree. Otherwise, chezmoi will run git pull --autostash --rebase [--\n" +
			"  recurse-submodules] , using chezmoi's builtin git if useBuiltinGit is true or if\n" +
			"  git.command cannot be found in $PATH.",
		example: "" +
			"  chezmoi update",
		longFlags: chezmoiset.New(
			"apply",
			"exclude",
			"include",
			"init",
			"parent-dirs",
			"recurse-submodules",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"P",
			"a",
			"i",
			"r",
			"x",
		),
	},
	"upgrade": {
		longHelp: "" +
			"Description:\n" +
			"  Upgrade chezmoi by downloading and installing the latest released version. This\n" +
			"  will call the GitHub API to determine if there is a new version of chezmoi\n" +
			"  available, and if so, download and attempt to install it in the same way as\n" +
			"  chezmoi was previously installed.\n" +
			"\n" +
			"  If the any of the $CHEZMOI_GITHUB_ACCESS_TOKEN, $CHEZMOI_GITHUB_TOKEN,\n" +
			"  $GITHUB_ACCESS_TOKEN, or $GITHUB_TOKEN environment variables are set, then the\n" +
			"  first value found will be used to authenticate requests to the GitHub API,\n" +
			"  otherwise unauthenticated requests are used which are subject to stricter rate\n" +
			"  limiting. Unauthenticated requests should be sufficient for most cases.",
		longFlags: chezmoiset.New(
			"executable",
			"method",
		),
	},
	"verify": {
		longHelp: "" +
			"Description:\n" +
			"  Verify that all targets match their target state. chezmoi exits with code 0\n" +
			"  (success) if all targets match their target state, or 1 (failure) otherwise. If\n" +
			"  no targets are specified then all targets are checked.",
		example: "" +
			"  chezmoi verify\n" +
			"  chezmoi verify ~/.bashrc",
		longFlags: chezmoiset.New(
			"exclude",
			"include",
			"init",
			"parent-dirs",
			"recursive",
		),
		shortFlags: chezmoiset.New(
			"P",
			"i",
			"r",
			"x",
		),
	},
}
