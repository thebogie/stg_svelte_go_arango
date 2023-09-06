/* eslint-disable global-require */
module.exports = {
    root: true,
    parser: '@typescript-eslint/parser',
    extends: [
        'eslint:recommended',
        'plugin:@typescript-eslint/recommended',
        'airbnb-base',
        'airbnb-typescript/base',
        'prettier',
    ],
    plugins: ['svelte3', '@typescript-eslint', 'import'],
    ignorePatterns: ['*.cjs', '*.mjs'],
    overrides: [
        {
            files: ['*.svelte'],
            processor: 'svelte3/svelte3',
            rules: {
                'import/no-extraneous-dependencies': 0,
                'import/extensions': 0,
                'import/newline-after-import': 0,
                'import/first': 0,
                'import/no-duplicates': 0,
                'import/no-mutable-exports': 0,
                'import/no-unresolved': 0,
                'import/prefer-default-export': 0,
                'import/order': 0,
                '@typescript-eslint/no-use-before-define': 0,
            },
        },
    ],
    settings: {
        'svelte3/typescript': () => require('typescript'),
        'import/resolver': {
            typescript: {
                alwaysTryTypes: true, // always try to resolve types under `<root>@types` directory even it doesn't contain any source code, like `@types/unist`
            },
        },
    },
    parserOptions: {
        sourceType: 'module',
        ecmaVersion: "latest",
        project: ['./tsconfig.json'],
        extraFileExtensions: ['.svelte'],
    },
    env: {
        browser: true,
        es2017: true,
        node: true,
    },
    globals: {
        Email: 'readonly',
    },
    rules: {
        '@typescript-eslint/no-shadow': 0,
        '@typescript-eslint/no-var-requires': 0,
        '@typescript-eslint/no-throw-literal': 0,
        'no-param-reassign': 0,
        'no-plusplus': 0,
        'consistent-return': 0,
        'no-lonely-if': 0,
        'prefer-const': 0,
        'import/no-extraneous-dependencies': 0,
        'import/prefer-default-export': 0,
    },
}