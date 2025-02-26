# Broadband GitHub Wiki

Welcome to the Broadband GitHub Wiki. For a rationale as to why this exists, see [here](https://wiki.uw.systems/posts/discovery-git-hub-wikis-j0ta1hcf).

**NOTE** - any changes made here in the GUI manually will be overwritten; this Wiki is automatically created and synced via the `/docs/wiki` directory in the root of this repo. If you want to make any changes, do it via there.

## Wiki Restrictions

This Wiki is not as nice a solution as I'd like it to be, and so comes with a few restrictions:

* The GitHub Wiki will flatten any hierarchy you establish here. All markdown files will be rendered in the root of the Wiki with no respect for folder structure.

* Linking to other markdown documents must *omit* the `.md` suffix (e.g. `[Go home](./Home)` instead of `[Go home](./Hom.md)`). This is because anything with a file extension is treated as a `rawcontext.github` link on the Wiki, and so `Home.md` would link to the raw markdown file (taking you out of the Wiki).

* Any images/assets you want to link to must also follow this flat structure; e.g. *all* images need to go into the root level `/images`, all PDFs need to go into `/assets/pdfs`, etc. This is an extremely annoying side-effect of the folder flattening issue from above: as the root of this repo is different to that of the Wiki.

* Only ever use *relative* imports. Absolute imports will always break in the Wiki as the root of the Wiki and this project will differ. You can probably fix this with some fancy `sed` script (I had a go, and lack the skills to make it work).
