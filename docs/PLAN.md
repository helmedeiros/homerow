# Delivery plan

## Goal

A terminal typing tutor whose first module — the **home row** — is fully
playable: 23 lessons covering `f j`, space, `d k`, `s l`, `a ;`, `g h`, the
left and right hands on their own, and the whole row together.

## Lesson types

| Type        | What happens                                                      | Scored |
|-------------|-------------------------------------------------------------------|--------|
| Intro       | A few short pages of advice (posture, where the fingers rest).    | no     |
| New keys    | Each key is shown on a keyboard with the finger that presses it, then drilled. | no (complete) |
| Review      | Mixed drill of the keys learned so far.                           | stars  |
| Practice    | Timed drill; speed matters more.                                  | stars  |
| Hand drill  | Only one hand's keys.                                             | stars  |
| Play        | Falling-letters game: clear letters before they land.             | stars  |

Stars: accuracy below 80% fails the exercise; 80/88/94/97% earn 1–4 stars, and
the fifth star needs 97% accuracy *and* the lesson's speed goal.

## Module 1 lesson map

| #  | Lesson                 | Type       | Keys            |
|----|------------------------|------------|-----------------|
| 1  | Welcome to touch typing| Intro      |                 |
| 2  | Keys f and j           | New keys   | f j             |
| 3  | Space bar              | New keys   | space           |
| 4  | Review: f and j        | Review     | f j             |
| 5  | Keys d and k           | New keys   | d k             |
| 6  | Review: d and k        | Review     | f j d k         |
| 7  | Practice: d and k      | Practice   | f j d k         |
| 8  | Play: f j d k          | Play       | f j d k         |
| 9  | Keys s and l           | New keys   | s l             |
| 10 | Review: s and l        | Review     | f j d k s l     |
| 11 | Practice: s and l      | Practice   | f j d k s l     |
| 12 | Keys a and ;           | New keys   | a ;             |
| 13 | Review: a and ;        | Review     | home 8          |
| 14 | First eight keys       | Practice   | home 8 (words)  |
| 15 | Play: first eight keys | Play       | home 8          |
| 16 | Finding home           | Intro      |                 |
| 17 | Home row: left hand    | Hand drill | a s d f g       |
| 18 | Home row: right hand   | Hand drill | h j k l ;       |
| 19 | Keys g and h           | New keys   | g h             |
| 20 | Review: g and h        | Review     | home row        |
| 21 | Practice: g and h      | Practice   | home row        |
| 22 | Home row review        | Practice   | home row (words)|
| 23 | Play: home row         | Play       | home row        |

## Delivery

Work is tracked in beads (`bd ready`). Each line below is one pull request made
of small commits; every commit passes `make check` locally and CI on GitHub
before the next one is pushed.

1. **Bootstrap** — module, Makefile, lint config, CI.
2. **Keyboard** — which finger and hand press each key.
3. **Typing session + scoring** — per-character marks, accuracy, words per
   minute, stars.
4. **Exercise generation** — seeded drills: repetitions, random groups, real
   words built only from learned keys.
5. **Curriculum** — lesson model and the 23 home row lessons.
6. **Application + storage** — use cases, ports, JSON progress file.
7. **Game** — falling-letters rules, deterministic and tick-driven.
8. **Terminal UI** — lesson map, intro pages, key introduction, typing screen
   with hands and live stats, results, game.
9. **CLI + end-to-end** — wiring, an end-to-end test that plays all 23 lessons,
   and a manual run in a real terminal.

## Testing

A balanced pyramid:

- **Unit** (most tests): every domain package, no I/O.
- **Application**: use cases against in-memory fakes of the ports.
- **Adapter**: the JSON store on a temp directory, sharing a contract test with
  the in-memory fake; UI models driven by key messages.
- **End-to-end**: the real program driven through every lesson of module 1, and
  a smoke test of the built binary.
