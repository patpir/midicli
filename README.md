# MIDI-CLI

Command-line interface for [miditf](https://github.com/patpir/miditf)

## Development

Build with `go build`.

## Usage

`midicli` operates on pipeline definitions.
These are stores in files, which are managed by various `midicli` commands.

To get started with `midicli`, you need to first initialize a new pipeline
definition.
`midicli init` creates a new pipeline definition file `MidiPipeline.json` in
the current working directory.
Use the `-f` flag to specify another filename (e.g.
`midicli init -f example.json`).

You can then add sources, transformations and visualizations to the newly
initialized pipeline definition.
The `midicli list` command enumerates all available sources, transformations
and visualizations - this is a *catalogue* of all block types.
Pass `sources`, `transformations` or `visualizations` to only list entries of
one category (e.g. `midicli list sources`).

To add an entry to the pipeline definition, use the `midicli add` command.
It takes the form of `midicli add <category> <type> <name> <arguments> ...`.
To read notes from a MIDI file named `music.midi`, use the command
`midicli add source midi-file import-music file=music.midi`, where `midi-file`
is the name of the source as shown by `midicli list`, `import-music` is a
user-chosen identifier unique to the pipeline definition, and `file` is the
name of the argument to the `midi-file` source, as shown by `midicli list`.

You can also remove existing items from the pipeline definition, by using the
`midicli remove` command with the name of the item to remove, e.g.
`midicli remove source import-music` to undo the add command in the example
above.

To show all entries in a pipeline definition, use the `midicli show` command.

If you want to clear a pipeline definition and return to the state created by
`midicli init`, use `midicli reset`.
This will remove all sources, transformations and visualizations from the
pipeline.

