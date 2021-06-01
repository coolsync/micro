// 1. define
// enum Option<T> {
//     Some(T),
//     None,
// }

// 2. use method
fn main() {
    let some_num = Some(5);

    let some_str = Some("a string");

    let absent_num: Option<i32> = None;

    let x = 1;
    let y: Option<i32> = Some(5);

    let mut tmp = 0;
    match y {
        Some(i) => tmp = i,
        None => {
            println!("no the value")
        }
    }
    let sum = x + tmp;
    println!("sum: {}", sum);

    // let r = plus_one(y);
    // match r {
    //     None => println!("nothing"),
    //     Some(i) => println!("r: {}", i)
    // }

    if let Some(value) = plus_one(y) {
        println!("value = {}", value);
    } else {
        println!("do nothing");
    }
}

fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(x) => Some(x + 1),
    }
}
