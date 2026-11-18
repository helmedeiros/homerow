# ADR 0002: Lesson kinds and games plug into the course

## Status

Accepted

## Context

The course grew from one module of drills and one game to six modules with
travel drills, tab practice, dynamic practice, capitals, numbers, symbols and
three games. Each addition risked spreading special cases through the UI, the
use cases and the content.

## Decision

A lesson is data, and its kind decides the steps it turns into.

- `curriculum.Lesson` holds a `Kind`, the allowed `Keys`, an optional `Focus`
  and a `Game`. `Validate` rejects lessons that cannot be played, such as a
  travel lesson without a home key or a capital without its small letter.
- `Lesson.Steps` turns a lesson into a list of `Step`s: pages, key intros,
  drills and games. Each kind picks its text from `exercise`: groups, words,
  capitalised words, numbers with words, symbols around words, travel or tabs.
- A game step carries its `Game`. The TUI maps each game to a screen, and every
  game reports the same `scoring.Result`, so results, stars and progress do not
  know which game was played. The pop-up warm-up ends without a score.
- Dynamic practice is the one kind that needs the player's history. The start
  use case loads the progress board and passes the most-missed keys to
  `Lesson.DynamicSteps`, keeping the curriculum free of storage.
- Content modules are plain lists of lessons built with small helpers
  (`newKeys`, `review`, `practice`, `travel`, `balloons`, `popUps`), and the
  catalog serves them in order.

## Consequences

- A new lesson kind is a `Kind`, a branch in `Lesson.Steps` and, when it needs
  new text, an `exercise` function. The UI and storage stay untouched.
- A new game is a domain type that reports a `scoring.Result` and a screen
  chosen in one place.
- Content stays declarative and is checked by `Validate`, so a lesson that
  cannot be played fails a test instead of reaching the player.
