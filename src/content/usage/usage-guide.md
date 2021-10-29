---
title: "Usage guide"
date: 2021-10-28T11:28:33-03:00
---
{{< toc >}}
You can have the full usage help by typing `todonotion -h` or for a specific command typing `todonotion <command> -h`

# Creating cards

## Creating a new card

This command form will create a card.

    todonotion new "Yay that is my first card"

As no lane and category were passed, the default lane is **To Do** and the default category is `todonotion`.

## Creating a new card in a specific lane

You can specify the destination lane by passing the `-l` flag:

    todonotion new "Yay that is my first card" -l "Doing"

## Creating a new tagged card

Tags are very usefull to organize and filter the cards, **todonotion** uses a property called `Property` (sorry for the poor imagination), that is a Select, enabling Notion to group the cards by their Property values.

    todonotion new "Yay that is my first card" -p happy

The card will have a Property with `happy` value:

[![card-1](../../images/card1.png)](../../images/card1.png)

# Moving cards

## Moving a single card

You can easly move a card between lanes with the `mv` command:


    todonotion mv "Yay that is my first card" -l Doing
    # or
    todonotion mv Yay -l Doing

The parameter can be the full or a partial card name.

{{< hint warning >}}
**Be careful**, the `mv` command finds the first card that name matchs and then moves it. Using partial names may have side effects if you have too many cards, causing confusion like matching/moving wrong cards. So if you have so many cards, please use a full (or almost full) card name.
{{< /hint >}}

## Moving many cards by tag

One more advantage to have tagged cards is that you can move all of them simultaneously by passing a `Property` value.
Using the `mvp` command you can move many cards by property:

    todonotion mvp happy -f "To Do" -t "Done"

As you can see, -f and -t are --from --to. So all cards with `happy` tag will move from "To Do" to "Done"

{{< hint info >}}
Here you need to specify the lanes to avoid moving done cards for example
{{< /hint >}}

# Listing cards

`todonotion` offers a powerful `ls` command that allows you to filter by **lane**, **property** and **title**. Also, you can combine all flags, having a powerful filter.

## Listing all cards in a specific lane

This command will list all Doing cards:

    todonotion ls -l Doing


## Listing all cards by tag

This command will list all cards that have this specifig Property value in "To Do"

    todonotion ls -p happy


## Listing all cards by title matching

This command will list all cards in "To Do" that have the word "task" in their titles.

    todonotion ls -t task


## Combining the flags

you can combine the flags to empower your `ls` command, like a filter.

    todonotion ls -t task -l Done -p happy

This command will list **only** the Done cards tagged as `happy` and that have **task** in some part of their titles.
