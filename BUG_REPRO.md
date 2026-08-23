# Bug Reproduction

## Bug

Rule evaluation and filtered rule-set results shared mutable slice backing arrays and tag maps with their callers. Sorting or appending a returned rule set could therefore change the original configuration and affect later reads.

## Trigger

Evaluate a rule slice that is retained by a revision, clone a rule with tags, or filter a rule set. Then mutate the returned value or evaluate it again. Before the fix, the original order, tags, or later revision contents changed across calls.

## Observed Error

The issue manifested as silent state contamination rather than a panic: repeated reads returned reordered rules or modified tags after an earlier simulation or filtered-result mutation.
