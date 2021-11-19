# PSOne

![psone_icon](./Playstation_icon.png)

## Introduction

Are you a **Veteran Unix Admin**? If so, you probably know the charm of the `PS1` environment variable.

For a deep focus I suggest you to take a look at this: https://en.wikipedia.org/wiki/Command-line_interface#Command_prompt.

## Why

The main reason why I decided to implement this utility is my fetishism for `PS1`. I like to customize this and change it for the right occasion, like to wear an elegant dress for your birthday 😄.

## How

The `psone` utility help you to manage different `PS1` storing them to a config file (`~/.psone.yaml`).

Whenever you want, you can **add**, **remove**, **set**, or **get** them. Once you set your favourite `PS1`, to apply your modification, just `source ~/.bashrc`.

## Examples

Fist of all, you need to generate your own `~/.psone.yaml` file. Here's an example of it:

```yaml
psones:
  default:
    value: "[\u@\h \W]\$ "
  halloween:
    value: "[\u@\h] 🎃👻🦇 > \[$(tput sgr0)\]"
  christmas:
    value: "[\u@\h] 🎅🎄❄️ \[\]"
```

Once you got your config file, you are ready to set your custom `PS1`. 

```bash
psone set halloween && source ~/.bashrc
```

This is going to be your result: `[user@laptop] 🎃👻🦇 > `.

Then, if you want to add a new one (the easter one):

```bash
psone add "easter" "[\u@\h] 🐇🥚 \[\]"
```

Here's the new updated `PS1`s list:

```bash
psone get
```

```yaml
psones:
  default:
    value: "[\u@\h \W]\$ "
  easter:
    value: "[\u@\h] 🐇🥚 \[\]"
  halloween:
    value: "[\u@\h] 🎃👻🦇 > \[$(tput sgr0)\]"
  christmas:
    value: "[\u@\h] 🎅🎄❄️ \[\]"
```

That's it.

For more information, just type: `psone help`.

## Installation

Just type the following command:

```bash
make
```

## Contributions

Of course, contributions are very welcome!

