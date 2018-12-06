module.exports = {
    extends: ["airbnb", "plugin:flowtype/recommended", "prettier"],
    plugins: ["flowtype", "prettier"],
    env: {
        browser: true,
        es6: true
    },
    rules: {
        "react/jsx-filename-extension": [1, { extensions: [".js", ".jsx"] }],
        "react/react-in-jsx-scope": 0,
        "import/extensions": 0,
        "react/prop-types": 0,
        "no-unused-vars": ["error", { varsIgnorePattern: "m", args: "none" }],
        "react/style-prop-object": 0,
        "prettier/prettier": "error",
        "react/style-prop-object": 0,
        "react/default-props-match-prop-types": 0,
        "react/no-unknown-property": 0,
        "no-nested-ternary": 0,
        "import/prefer-default-export": 0, // TODO: remove
        "react/jsx-closing-bracket-location": 0,
        "react/jsx-closing-tag-location": 0,
        "react/jsx-curly-spacing": 0,
        "react/jsx-equals-spacing": 0,
        "react/jsx-first-prop-new-line": 0,
        "react/jsx-indent": 0,
        "react/jsx-indent-props": 0,
        "react/jsx-max-props-per-line": 0,
        "react/jsx-one-expression-per-line": 0,
        "react/jsx-tag-spacing": 0,
        "react/jsx-wrap-multilines": 0,
        "jsx-a11y/label-has-for": 'off', //TODO: remove
        "jsx-a11y/label-has-associated-control": 'off', //TODO: remove
        "jsx-a11y/anchor-is-valid": 0, //TODO: remove
        "jsx-a11y/alt-text": 0, //TODO: remove
        "react/button-has-type": 0, //remove
        "jsx-a11y/no-static-element-interaction": 0, //remove
        "jsx-a11y/click-events-have-key-events": 0, //remove
        "jsx-a11y/no-noninteractive-element-interactions": 0, //remove
        "array-callback-return": 0, //remove(maybe)
        "no-return-assign": 0,
        camelcase: 0,
        "import/no-mutable-exports": 0,
        "no-console": 0, //remove
        "no-var": 0, // remove after research
        "jsx-a11y/no-static-element-interactions": 0, //remove
        "react/no-array-index-key": 0, // research and remove
        "react/no-unescaped-entities": 0, //remove
        "jsx-a11y/anchor-has-content": 0,
        "flowtype/no-types-missing-file-annotation": 0, //remove after research
        "no-undef": 0, //important remove
        "react/no-string-refs": 0,
        "one-var": 0,
        "react/no-this-in-sfc": 0,
        "consistent-return": 0,
        "no-plusplus": 0,
        "vars-on-top": 0,
        "no-shadow": 0,
        "no-param-reassign": 0,
        "new-cap": 0,
        "no-prototype-builtins": 0,
        "import/no-cycle": 0,
        "prefer-destructuring": ["error", { object: true, array: false }] // remove later
    },
    settings: {
        "import/resolver": {
            node: {
                moduleDirectory: ["node_modules", "src"]
            }
        }
    },
    parser: "babel-eslint",
    parserOptions: {
        allowImportExportEverywhere: true,
        ecmaVersion: 6,
        sourceType: "module",
        ecmaFeatures: {
            jsx: true
        }
    }
};
