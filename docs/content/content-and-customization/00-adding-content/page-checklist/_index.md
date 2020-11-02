---
date: "2017-04-24T18:36:24+02:00"
description: ""
title: As a Checklist
weight: 190
---

A basic .md file can be rendered as a form/checklist/questionnaire.

{{% notice %}}
**A page rendered as a Checklist** is a page with a special rendered TOC, and a LOAD/DOWNLAOD form results buttons.
\
[{{%icon aspect_ratio%}} click here to view an example]({{%ref "checklist/_index.md"%}})

{{%/notice%}}

To tell Hugo to consider a page as a checklist, just add a `checklist: true` in the frontmatter of your page.

```yaml
---
checklist: true
---
```

## Demo
[{{%icon aspect_ratio%}} click here to view an example]({{%ref "checklist/_index.md"%}})

## Dedicated shortcodes
* c/text - a line or box input type
* c/choices - a radiogroup or a checkboxgroup
* c/list - a select or a multiselect
* c/check - a checkbox
* c/switch - a switch
* c/show - display a part with a condition using user inputs
* c/hidden - set a hidden value
