
## Requirements

To make it easier for us to work on different platforms and configurations we use containers to achieve a unified development environment.
1. Install [Docker](https://docs.docker.com/install/).
2. Install [Docker Compose](https://docs.docker.com/compose/install/).
3. [Atom](https://atom.io/) text editor.

## Quick Start

For a **quick start** run:

    $ docker-compose -f docker-compose.dev.yml up;

If your setup is correct, the development server URL will be copied to your clipboard and the project should be served on that URL

For more flexibility and the ability to run custom commands on the container run:

    $ docker-compose -f docker-compose.dev.yml run --service-ports --rm app /bin/bash

## Commands

- `$ yarn app:hot` Starts development environment with hot/live reloading.
- `$ app:watch` Starts development environment watching source files, without hot/live reloading.
- `$ app:build` Builds production version.

## Coding Standards and Conventions

1.  **Indentation**: Use _tabs_ not _spaces_.
2.  **Spacing**: Use 2 spaces per tab in .js files.
3.  **File naming**: Use lowercase with dashes (-).
4.  **Variable naming**: use _lowercase_ and _underscores_ eg. `let var_one = 1;`
5.  **Constants**: use uppercase eg. `const PI = 3.142;`
6.  **File Headers**: On save modify file headers or install Atom `file-heading` plugin.
7.  **Testing**: Write tests for every feature and
    functionality you add. Code without tests will not be merged.
8.  **Pushes and Commits**: Create a **feature branch** for every **commit** you make.
9.  **Commit Messages**: Let your commit message be as detailed as possible. Ambiguous commits will not be merged. 
10. **Code Comments**: Comment, comment and comment everything. Include todo, fixme and other annotations where necessary.
11. **DocBlocks**: Add [**docblocks**](https://atom.io/packages/docblockr) where necessary, they are very important.

## Contributors

1.  [Princewill Samuel](mailto:prinzllsamuel@gmail.com)
2.  [Israel Duff](mailto:theorix8@gmail.com)
3.  [Micheal Eneji](mailto:bigmikeeneji@gmail.com)
4.  [Thomas Agba](mailto:agbatom@gmail.com)
5.  [Inim Andrew](mailto:grapheneee@gmail.com)