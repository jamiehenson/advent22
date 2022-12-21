import math
import os


class Operation:
    def __init__(self, constant=0, operator="+"):
        self.constant = constant
        self.operator = operator


class Monkey:
    def __init__(
        self,
        id=-1,
        items=[],
        operation=Operation(),
        testDivisor=1,
        testTrueRecipient=-1,
        testFalseRecipient=-1,
    ):
        self.id = id
        self.items = items
        self.operation = operation
        self.testDivisor = testDivisor
        self.testTrueRecipient = testTrueRecipient
        self.testFalseRecipient = testFalseRecipient


def generateMonkeys(src):
    monkeys = {}

    startingMonkeys = src.split("\n\n")
    for monkeyText in startingMonkeys:
        lines = monkeyText.split("\n")
        monkey = Monkey()

        for line in lines:
            splitLine = line.split(" ")
            last = splitLine[-1]

            if "Monkey" in line:
                monkey.id = int(last.replace(":", ""))
            elif "Starting items" in line:
                monkey.items = list(map(int, line.split(": ")[1].split(", ")))
            elif "Operation" in line:
                monkey.operation = Operation(
                    constant=splitLine[-1], operator=splitLine[-2]
                )
            elif "Test" in line:
                monkey.testDivisor = int(last)
            elif "If true" in line:
                monkey.testTrueRecipient = int(last)
            elif "If false" in line:
                monkey.testFalseRecipient = int(last)

        monkeys[monkey.id] = monkey

    return monkeys


def main(turns, verbose=False):
    path = os.path.join(os.path.dirname(__file__), "monkeys.txt")
    src = open(path, "r").read()
    monkeys = generateMonkeys(src)

    monkeyInspections = {}
    for key in monkeys:
        monkeyInspections[key] = 0

    for i in range(turns):
        for monkeyKey in monkeys:
            if verbose:
                print(f"\nMonkey {monkeyKey}:")
            monkey = monkeys[monkeyKey]
            for item in monkey.items:
                id = monkey.id
                operator = monkey.operation.operator
                constant = monkey.operation.constant
                testDivisor = monkey.testDivisor
                testTrueRecipient = monkey.testTrueRecipient
                testFalseRecipient = monkey.testFalseRecipient

                if verbose:
                    print(f"\tMonkey inspects an item with a worry level of {item}")
                monkeyInspections[id] += 1

                # Monkey handles the item - worry increases
                newItem = item

                constant = int(item if constant == "old" else constant)

                if operator == "+":
                    newItem += constant
                    if verbose:
                        print(f"\t\tWorry level increases by {constant} to {newItem}.")
                elif operator == "*":
                    newItem *= constant
                    if verbose:
                        print(
                            f"\t\tWorry level is multiplied by {constant} to {newItem}."
                        )

                # Monkey gets bored - worry decreases
                newItem = newItem // 3
                if verbose:
                    print(
                        f"\t\tMonkey gets bored with item. Worry level is divided by 3 to {newItem}."
                    )

                # Worry conditionally divides
                itemRecipient = id
                if newItem % testDivisor == 0:
                    if verbose:
                        print(f"\t\tCurrent worry level is divisible by {testDivisor}.")
                    itemRecipient = testTrueRecipient
                else:
                    if verbose:
                        print(
                            f"\t\tCurrent worry level is not divisible by {testDivisor}."
                        )
                    itemRecipient = testFalseRecipient
                if verbose:
                    print(
                        f"\t\tItem with worry level {newItem} is thrown to monkey {itemRecipient}"
                    )

                # Pass the item off
                monkeys[itemRecipient].items.append(newItem)

            # Wipe all items on monkey - all should have been passed off by this point
            monkeys[id].items = []

        print(f"\nAfter Round {i+1}:")
        for monkeyKey in monkeys:
            print(f"Monkey {monkeyKey}: {monkeys[monkeyKey].items}")

    print("\nMonkey inspections:")
    for key in monkeyInspections:
        quantity = monkeyInspections[key]
        inflected = "time" if quantity == 1 else "times"
        print(f"Monkey {key} inspected items {quantity} {inflected}")

    sortedInspections = list(
        dict(
            sorted(monkeyInspections.items(), key=lambda item: item[1], reverse=True)
        ).values()
    )

    print(f"\nTotal monkey shenanigans: {sortedInspections[0] * sortedInspections[1]}")


main(20, verbose=False)
