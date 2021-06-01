//1、创建一个空String
//2、通过字面值创建一个String
//2.1、使用String::from()
//2.2、使用str的方式
//3、更新String
//3.1、push_str
//3.2、push
//3.3、使用“+”合并字符串
//3.4、使用format!
//4、String 索引
//5、str 索引
//6、遍历
//6.1、chars
//6.2、bytes


fn main() {
    // 1 
    let mut s0 = String::new();
    s0.push_str("s0 hello");
    println!("s0: {}", s0);

    // 2
    let s1 = String::from("init some string s1");
    println!("s1: {}", s1);

    let s1 = "init some string s1".to_string();
    println!("s1: {}", s1);

    // 3
    let mut s1 = String::from("hello");
    s1.push_str(" world");
    let ss = "!".to_string();
    s1.push_str(&ss);
    println!("s1: {}", s1);
    println!("ss: {}", ss);

    let mut s2 = String::from("tea");
    s2.push('m');
    println!("s2 = {}",  s2);

    let s1 = "hello".to_string();
    let s2 = String::from(" world");
    let s3 = s1 + &s2;  // s1 ownership move to s3; + method
    // println!("s1: {}", s1); // borrow of moved value: `s1`
    println!("s2: {}", s2);
    println!("s3: {}", s3);

    // format!
    let s11 = String::from("aaa");
    let s22 = String::from("bbb");
    let s33 = String::from("ccc");
    let s = format!("{}-{}-{}", s11, s22, s33);  // like pirntln!
    format!("{}", s22);
    format!("{}", s33);

    // String index
    let s1 = String::from("hello");
    // let h = s1[0]; // error
    
    // str index
    let s2 = "你好";
    let h1 = &s2[0..3];
    println!("h1: {}", h1);

    // iterate
    for c in s2.chars(){
        println!("c = {}", c);
    }
    for b in s2.bytes() {
        println!("b = {}", b);
    }
    println!("Hello, world!");




}
