module.exports = {
    rules: {
        "scope-case": [2, "always", "lower-case"],
        "subject-case": [2, "always", "lower-case"],
        "subject-empty": [2, "never"],
        "subject-full-stop": [2, "never", "."],
        "subject-max-length": [2, "always", "40"],
        "subject-min-length": [2, "always", "5"],
        "type-enum": [2, "always", ['feat', 'fix', 'docs', 'style', 'refactor', 'test', 'revert']],
        "type-case": [2, "always", "lower-case"],
    }
}
