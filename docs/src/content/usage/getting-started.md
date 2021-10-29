---
title: "Getting Started"
date: 2021-10-27T11:27:37-03:00
---

{{< toc >}}

# Introduction

Wellcome to **todonotion** Get Started page!

**todonotion** is a terminal based app that uses the **Notion.so Task List** to replicate the Kanban board
for **To Do**, **Doing** and **Done**.

You can create, list, move and delete cards without leaving the terminal! And using the Notion.so page as a database, which is the core and the bonus of our utilitary!

## Installation

{{< hint ok >}}
You can find all detais in our installation guide on {{< button relref="installation/" >}}Installation{{< /button >}} page. Please go there!
{{< /hint >}}


## Configuring the access to Notion API

First of all, duplicate this [task list template](https://www.notion.so/9734e8c4df1441f594e44d9c068c7780?v=005250cb1ef74bacb2e466530802510e) into your notion workspace, we will use it to speed up our setup.

{{< hint info >}}
The ***duplicate*** button may appear in top-right corner.
{{< /hint >}}


We also need the following environment variables:

* NOTION_KEY
* NOTION_DATABASE_ID
  

Follow the **step1** [this guide](https://developers.notion.com/docs/getting-started) to claim your **NOTION_KEY** and follows the **step2**, but instead of creating some new database, uses the database that you create from our duplicated template.

Now that you have them, export the variables:

    export NOTION_KEY=<your secret>
    export NOTION_DATABASE_ID=<your database created above>


That's enough to get `todonotion` working like a charm.
