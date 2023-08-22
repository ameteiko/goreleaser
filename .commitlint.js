module.exports = {
    rules: {
        "scope-case": [2, "always", "lower-case"],
        // "subject-case": [2, "always", ['lower-case', 'sentence-case', 'start-case', 'pascal-case', 'upper-case']],
        "subject-empty": [2, "never"],
        "subject-full-stop": [2, "never", "."],
        "subject-max-length": [2, "always", "120"],
        "subject-min-length": [2, "always", "5"],
        "type-enum": [2, "always", ['feat', 'fix', 'docs', 'style', 'refactor', 'test', 'revert']],
        "type-case": [2, "always", "lower-case"],
    }
}
