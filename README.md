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

An *array* of "exec" ought to be present in the configuration file, each
containing an *array* of commands and parameters to run.

The order in which the items appear in the configuration file will be the order
in which rofi will display the options.

Example configuration (YAML):

    ---
    - name: Laptop Mode
      exec:
        -
          - notify-send
          - laptop mode
          - Enabling...
        -
          - laptop-mode
          - enable
        -
          - notify-send
          - laptop mode
          - Enabled
    - name: Shut down
      exec: 
        -
          - shutdown
          - -h
          - now
    - name: Reboot (sudo)
      exec:
        -
          - sudo
          - shutdown
          - -r
          - now

## Copyright and License

`rofixec` is Copyright (c) 2021, Marco Fontani <MFONTANI@cpan.org>

It is released under the MIT license - see the `LICENSE` file in this repository/directory.
