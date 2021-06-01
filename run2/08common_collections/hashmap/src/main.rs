use std::collections::HashMap;

fn main() {
    // 1. create a new hash map
    let mut scores = HashMap::new();

    scores.insert(String::from("blue"), 10);
    scores.insert(String::from("yellow"), 50);

    // 2. use tuple vector collect method create hash map
    let teams = vec![String::from("blue"), String::from("yellow")];
    let init_scores = vec![10, 50];

    let scores: HashMap<_,_> = teams.iter().zip(init_scores.iter()).collect();

    // 3. Hash map and ownership
    let field_name = String::from("like color");
    let field_color = String::from("blue");

    let mut map = HashMap::new();
    map.insert(&field_name, field_color);
    println!("field_name: {}", field_name);

    // 4. Access the value in the hash map
    let team_name = String::from("blue");
    let score = scores.get(&team_name);

    // println!("blue score: {:?}", score);
    if let Some(val) = score {
        println!("score val: {}", &val);
    }

    // traveser
    for (key, val) in &scores {
        println!("{}, {}", key, val);
    }

    // 5. update hash map
    // 覆盖一个值 overwrite a Value
    let mut scores = HashMap::new();
    scores.insert(String::from("blue"), 10);
    scores.insert(String::from("blue"), 50);
    
    println!("{:?}", scores);

    // Insert only when the key has no corresponding value
    let mut scores = HashMap::new();
    scores.insert(String::from("blue"), 10);
    
    scores.entry(String::from("yellow")).or_insert(50);
    scores.entry(String::from("blue")).or_insert(50);   // has, no change
    println!("{:?}", scores);

    // Update a value based on the old value
    let text = "hello world wonderful world";

    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }

    println!("{:?}", map);
    println!("Hello, world!");
}
