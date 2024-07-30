CREATE TABLE IF NOT EXISTS accounts
(
    account_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    current_value INTEGER NOT NULL DEFAULT 0,
    currency TEXT,
    rules INT
);
CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
    value INTEGER NOT NULL,
    date DATETIME,
    description TEXT
);
CREATE TABLE IF NOT EXISTS rules(
    rule_id INTEGER PRIMARY KEY AUTOINCREMENT,
    rule_name VARCHAR(255) NOT NULL,
    rule_description TEXT
);
CREATE TABLE IF NOT EXISTS rules_for_account(
    account_id INTEGER,
    rule_id INTEGER,
    state BOOLEAN,
    FOREIGN KEY (account_id) references accounts(account_id),
    FOREIGN KEY (rule_id) references rules(rule_id)
);