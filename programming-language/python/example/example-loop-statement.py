import random

# [Keyword] list: fixed tier labels and index labels.
# - tiers: each array must include one member from each tier a-e.
# - index_names: target arrays (a_index_array, b_index_array, c_index_array).
tiers = ["a", "b", "c", "d", "e"]
index_names = ["a", "b", "c"]

# [Keyword] dict type hint: dict[str, list[str]]
# Tier member pool. Each tier has exactly 3 unique members because we build 3 arrays.
tier_members: dict[str, list[str]] = {
    "a": ["a1", "a2", "a3"],
    "b": ["b1", "b2", "b3"],
    "c": ["c1", "c2", "c3"],
    "d": ["d1", "d2", "d3"],
    "e": ["e1", "e2", "e3"],
}

# [Keyword] dictionary comprehension + random.sample(..., k=...)
# For each tier, randomly reorder all members once, then distribute by position.
# This guarantees no duplicate member between arrays within the same tier.
distributed_by_tier: dict[str, list[str]] = {
    tier: random.sample(members, k=len(index_names))
    for tier, members in tier_members.items()
}

index_arrays: dict[str, list[str]] = {}
# [Keyword] enumerate: use idx to pick one unique member per tier for each array.
# [Keyword] list comprehension: collect one member from each tier.
# [Keyword] random.shuffle: randomize display/order inside each array.
for idx, name in enumerate(index_names):
    picked = [distributed_by_tier[tier][idx] for tier in tiers]
    random.shuffle(picked)
    index_arrays[f"{name}_index_array"] = picked

# [Keyword] flatten with list comprehension + set + assert
# Flatten all arrays into one list and verify global uniqueness.
# If duplicate exists, assert raises an error.
all_values = [value for values in index_arrays.values() for value in values]
assert len(all_values) == len(set(all_values)), "Global duplicate detected"

# Print final result.
for key, value in index_arrays.items():
    print(f"{key}: {value}")