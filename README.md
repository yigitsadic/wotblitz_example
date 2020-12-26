World of Tanks Blitz Tanks example project.

This is not official tanks database.

I have no relationship with World of Tanks or Wargaming.

This is just a hobby project which I give a try to entgo.

GraphQL Playground link:
https://guarded-savannah-12730.herokuapp.com


### Query Examples

Tank details
```graphql
query {
    tank(name: "AMX 38") {
        name
        isPremium
        tier
        country
        tankClass

        nextTanks {
            name
        }

        previousTanks {
            name
        }

        stockModules: modules(status: Stock) {
            type
            name
        }

        upgradeModules: modules(status: Upgrade) {
            type
            name
        }
    }
}
```

Search
```graphql
query {
  search(term: "tiger") {
    ... on Tank {
      name
      tier
      country
      isPremium
    }
    
    ... on Module {
      name
      type
    }
  }
}
```

Filter Tanks
```graphql
query {
  filterTanks(tier: 7, tankClass: TankDestroyer, isPremium: false) {
    name
    country
  }
}
```

Get tanks of tech tree
```graphql
query {
    techTree(country: France) {
        name
        isPremium
        tier
        tankClass
        modules {
            name
            type
        }

        nextTanks {
            name
        }

        previousTanks {
            name
        }
    }
}
```
