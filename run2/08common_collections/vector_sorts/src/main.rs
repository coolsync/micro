fn main() {
    let v1 = vec![100, 32, 57, 101, 33, 58, 1, 3, 5];

    let v2: Vec<i32> = Vec::new();

    let tmp = 0;
    for val in &v1 {
        println!("val: {}", val);
    }
    println!("Hello, world!");
}
