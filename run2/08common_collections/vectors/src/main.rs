// 1. create a new, empty vector
// 2. create has init value vector
// 3. drop vector
// 4. get elem
// 5. update elem
// 6. traveser
// 7. enum type vector
fn main() {
    // 1. create a new, empty vector
    // let v: Vec<i32> = Vec::new();
    let mut v: Vec<i32> = Vec::new();
    // v.push(1);

    // 2. create has init value vector
    let v = vec![1, 2, 3];

    // 3. drop vector
    {
        let v = vec![1, 2, 3];
    } // leave scope

    // 4. get elem
    let one: &i32 = &v[0];
    // let four: &i32 = &v[3];
    println!("one = {}", one);
    println!("one = {}", *one);

    //(2)推荐的方法
    // match v.get(1) {
    match v.get(3) {
        Some(value) => println!("value = {}", value),
        None => println!("None"),
    }

    // 5. update elem
    // let v2 = vec![1,2,3];
    let mut v2: Vec<i32> = Vec::new();
    v2.push(1);
    v2.push(2);
    v2.push(3);

    // 6
    // 不可变 traveser
    for i in &v2 {
        println!("i = {}", i)
    }
    // 可变 traveser
    for i in &mut v2 {
        *i += 1;
        println!("i: {}", i)
    }

    // 7
    enum Context {
        Text(String),
        Int(i32),
        Float(f32),
    }

    let v3 = vec![
        Context::Text(String::from("hello")),
        Context::Int(-1),
        Context::Float(0.02),
    ];

    // 8 补充
    let mut v3 = vec![1,3,5];
    let first = &v3[0]; // -- immutable borrow occurs here
    // v3.push(7); // mutable borrow occurs here, v3 rechange
    println!("first: {}", first);


    println!("Hello, world!");
}
