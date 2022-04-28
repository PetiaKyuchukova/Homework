

CREATE TABLE articles(
  id   INT  NOT NULL  PRIMARY KEY,
  title text  NOT NULL,
  score int  NOT NULL,
  time text  NOT NULL
);

-- name: CreateArticle :execresult
INSERT INTO articles (
  id, title, score, time
) VALUES (
  ?, ?, ?, ?
);

-- name: DeleteArticles :exec
DELETE FROM articles;

-- name: GetArticles :many
SELECT * FROM articles;