# Delivery plan

## Goal

A terminal typing tutor whose first module — the **home row** — is fully
playable: 23 lessons covering `f j`, space, `d k`, `s l`, `a ;`, `g h`, the
left and right hands on their own, and the whole row together.

Module 1 is delivered. The course then grows to six modules: top row, bottom
row, shift key, numbers and symbols, 187 lessons in all (see below).

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

## Modules 2 to 6

The course grows to six modules and 187 lessons. Lesson numbers run on
without gaps: modules 2 and 3 keep the numbers of the reference course
(24–88); shift, numbers and symbols are renumbered to follow on (89–187).

| Module | Lessons | New keys |
|--------|---------|----------|
| 1 Home Row   | 1–23    | a s d f g h j k l ; space |
| 2 Top Row    | 24–51   | r u e i w o q y p t |
| 3 Bottom Row | 52–88   | v m c , x . z / b n tab |
| 4 Shift Key  | 89–142  | capitals, : ? |
| 5 Numbers    | 143–157 | 1–0 |
| 6 Symbols    | 158–187 | $ & # * @ ( ! ) % ^ |

### New lesson types

| Type     | What happens |
|----------|--------------|
| Travel   | Reach from a home key to its keys above and below and back, e.g. `juj jmj jhj`. |
| Dynamic  | A drill weighted towards the keys you miss most, from per-key mistakes kept in the progress file. |
| Tab      | Practice with tab characters, shown as `→`. |

### Games

| Game | Used by | Scored |
|------|---------|--------|
| Falling letters (existing) | most play lessons | yes |
| Balloons: whole words rise; type the word to pop it | 38, 45, 62, 74, 110, 124, 132, 173 | yes |
| Pop-up: keys pop up and hide again unless typed in time | 50, 82, 103, 117 | no, a warm-up |

### Module 2: Top Row (24–51)

| # | Lesson | Type |
|---|--------|------|
| 24 | Keys r and u | New keys |
| 25 | Review: r and u | Review |
| 26 | Practice: r and u | Practice |
| 27 | Keys e and i | New keys |
| 28 | Review: e and i | Review |
| 29 | Practice: e and i | Practice |
| 30 | Sit straight | Intro |
| 31 | Review: home + ruei | Practice |
| 32 | Top row: left hand | Hand drill |
| 33 | Top row: right hand | Hand drill |
| 34 | Play: ruei | Play (falling) |
| 35 | Keys w and o | New keys |
| 36 | Review: w and o | Review |
| 37 | Practice: w and o | Practice |
| 38 | Play: top six keys | Play (balloons) |
| 39 | Keys q and y | New keys |
| 40 | Review: q and y | Review |
| 41 | Practice: q and y | Practice |
| 42 | Keys p and t | New keys |
| 43 | Review: p and t | Review |
| 44 | Practice: p and t | Practice |
| 45 | Play: top row | Play (balloons) |
| 46 | Think in words | Intro |
| 47 | Review: home + top | Practice |
| 48 | Top row: left hand | Hand drill |
| 49 | Top row: right hand | Hand drill |
| 50 | Play: top row | Play (pop-up) |
| 51 | Play: top row | Play (falling) |

### Module 3: Bottom Row (52–88)

| # | Lesson | Type |
|---|--------|------|
| 52–54 | Keys, review, practice: v and m | New keys, Review, Practice |
| 55, 56 | Travel right / left index | Travel |
| 57–59 | Keys, review, practice: c and , | New keys, Review, Practice |
| 60, 61 | Travel right / left middle | Travel |
| 62 | Play: v m c , | Play (balloons) |
| 63 | History of QWERTY | Intro |
| 64–66 | Keys, review, practice: x and . | New keys, Review, Practice |
| 67, 68 | Travel right / left ring | Travel |
| 69–71 | Keys, review, practice: z and / | New keys, Review, Practice |
| 72, 73 | Travel right / left pinky | Travel |
| 74 | Play: bottom row | Play (balloons) |
| 75–77 | Keys, review, practice: b and n | New keys, Review, Practice |
| 78 | Take a break | Intro |
| 79, 80 | Bottom row: left / right hand | Hand drill |
| 81 | Practice: bottom row | Practice |
| 82 | Play: bottom row | Play (pop-up) |
| 83 | Play: bottom row | Play (falling) |
| 84, 85 | Travel right / left, all fingers | Travel |
| 86 | One small space: the tab key | Intro |
| 87 | Practice: tab | Tab |
| 88 | Dynamic practice | Dynamic |

### Module 4: Shift Key (89–142)

Capitals are typed by holding the shift key of the *other* hand with its
pinky. The hands show both fingers. A terminal cannot tell which shift key was
pressed, so the rule is taught and shown but not enforced.

| # | Lessons |
|---|---------|
| 89 | Shift key intro |
| 90–92 | Capital F and J: keys, review, practice |
| 93–95 | Capital D and K |
| 96 | Play: F J D K (falling) |
| 97–99 | Capital S and L |
| 100–102 | Capital A and : |
| 103 | Play: S L A : (pop-up) |
| 104–106 | Capital G and H |
| 107–109 | Capital T and Y |
| 110 | Play: G H T Y (balloons) |
| 111–113 | Capital R and U |
| 114–116 | Capital E and I |
| 117 | Play: R U E I (pop-up) |
| 118–120 | Capital W and O |
| 121–123 | Capital Q and P |
| 124 | Play: W O Q P (balloons) |
| 125–127 | Capital V and M |
| 128 | Travel index |
| 129–131 | Capital C and ? |
| 132 | Play: V M C ? (balloons) |
| 133–135 | Capital Z and N |
| 136–138 | Capital X and B |
| 139–141 | Travel middle, ring, pinky |
| 142 | Dynamic practice |

### Module 5: Numbers (143–157)

Keys, review and practice for each pair: 4 and 7 (143–145), 3 and 8
(146–148), 2 and 9 (149–151), 1 and 0 (152–154), 5 and 6 (155–157).

### Module 6: Symbols (158–187)

Keys, review and practice for $ and & (158–160), # and * (161–163), @ and (
(164–166), ! and ) (167–169), % and ^ (170–172); play all symbols (173);
travel to each of $ # @ ! % & * ( ) ^ (174–183); travel left and right
(184–185); review (186) and practice (187) of all symbols.

### Cross-cutting work

- **Course:** the catalog serves several modules; the lesson list shows module
  headers, scrolls, and jumps between modules with ← →; `homerow lessons`
  groups by module.
- **Keyboard:** shifted characters know their key and which shift to hold;
  tab belongs to the left pinky.
- **Progress:** stays keyed by lesson number. Version 2 of the file adds
  mistakes per key for dynamic practice, and version 1 files still load.

### Delivery of modules 2 to 6

4–7 small commits each weekday evening, from 6 to 24 November 2026: one branch
and pull request per day, merged after the day's last commit. Every commit
passes `make check` and CI before the next push. Each module ends with a test
that plays all of its lessons against the file store, plus a run of the real
binary in a terminal.

| Day | Work |
|-----|------|
| Fri 6 Nov | course model, catalog, module headers |
| Mon 9 Nov | module switching, listing, shift-aware keyboard and hands |
| Tue 10 Nov | top row content 24–51 |
| Wed 11 Nov | balloons game |
| Thu 12 Nov | pop-up game, top row end to end |
| Fri 13 Nov | travel lessons, bottom row 52–74 |
| Mon 16 Nov | bottom row 75–87, tab, per-key mistakes |
| Tue 17 Nov | dynamic practice, bottom row end to end |
| Wed 18 Nov | capitals, shift 89–110 |
| Thu 19 Nov | shift 111–142, end to end |
| Fri 20 Nov | numbers 143–157 |
| Mon 23 Nov | symbols 158–187 |
| Tue 24 Nov | full course end to end, docs |
