World of Tanks Blitz Tanks example project.

This is not official tanks database.

I have no relationship with World of Tanks or Wargaming.

This is just a hobby project.

GraphQL Playground link:
https://guarded-savannah-12730.herokuapp.com


### Query Examples

Tank details
```graphql
query {
  tank(name: "R35") {
    name
    isPremium
    tier
    country
    tankClass

    stockModules: modules(isStock: true) {
      type
      name
    }

    upgrableModules: modules(isStock: false) {
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
    }
}
```
