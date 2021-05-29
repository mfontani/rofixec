# rofixec

The goal is to provide a single binary which can be configured to both present
a (user-sorted) list of actions to take, and take the user-specified action,
which can comprise of one or multiple command executions.

Example invocation:

    rofi -modi foo:'rofixec -config foo.yaml' -show foo
    rofi -modi foo:'rofixec -config foo.json' -show foo

Sample use-cases are provided in the examples/ directory.

The configuration can be either in YAML or JSON. The format used is the same
for YAML and JSON:

An *array* of "items" ought to be present in the configuration file.

The order in which the items appear in the configuration file will be the order
in which rofi will display the options.

Each "item" is required to have a `name` property. This is the piece of text
that will be shown to the user when they pick options. Be creative.

Each "item" can have both:

- an `exec` string property, and an optional `args` array of strings property,
  which will be always executed asynchronously once the user has picked the
  corresponding `name`, and
- an array of "items" under the `commands` property, each containing an `exec`
  property, and an optional array of `args`.

Example configuration (YAML):

    ---
    - name: Laptop Mode
      exec: notify-send
      args:
        - laptop mode
        - Enabling...
      commands:
        - exec: laptop-mode
          args:
            - enable
        - exec: notify-send
          args:
            - laptop mode
            - Enabled
    - name: Shut down
      exec: shutdown
      args:
        - -h
        - now
    - name: Reboot (sudo)
      exec: sudo
      args:
        - shutdown
        - -r
        - now

## Copyright and License

`rofixec` is Copyright (c) 2021, Marco Fontani <MFONTANI@cpan.org>

It is released under the MIT license - see the `LICENSE` file in this repository/directory.
