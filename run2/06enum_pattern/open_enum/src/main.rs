// 1 类c define
enum IpAddrKind {
    V4,
    V6,
}

struct IpAddr {
    kind: IpAddrKind,
    address: String,
}

//2、run 提倡的方式定义
enum IpAddr2 {
    V4(String),
    V6(String),
}

// 3、可以是 不同的type
enum IpAddr3 {
    V4(u8, u8, u8, u8),
    V6(String),
}


//4、经典用法
enum Message {
    Quit,
    Move{x: i32, y: i32},
    Write(String),
    Change(i32, i32, i32),
}

//5、 enum type mehod, match
impl Message {
    fn print(&self) {
        match *self {
            Message::Quit => println!("quit"),
            Message::Move{x, y} => println!("x : {}, y : {}", x, y),
            Message::Change(x, y, z) => println!("x = {}, y = {}, z = {}", x, y,z),
            // Message::Write(&s) => println!("s = {}", s),
            _ => println!("write")
        }
    }
}
// 等同于
// struct QuitMessage, 类单元结构体
// struct Move {x: i32, y: i32}
// struct Wirte(String), tuple struct
// 
fn main() {
    // 1
    let i1 = IpAddr {
        kind: IpAddrKind::V4,
        address: String::from("127.0.0.1"),
    };

    let i2 = IpAddr {
        kind: IpAddrKind::V6,
        address: String::from("::1"),
    };
    // 2
    let i1 = IpAddr2::V4(String::from("127.0.0.1"));
    let i2 = IpAddr2::V6(String::from("::1"));

    // 3
    let i1 = IpAddr3::V4(127, 0, 0, 1);
    let i2 = IpAddr3::V6(String::from("::1"));

    // 4 
    let q = Message::Quit;
    q.print();

    let mo = Message::Move{x:1, y:2};
    mo.print();
    
    let cg = Message::Change(1, 2, 3);
    cg.print();

    let wri = Message::Write(String::from("hello"));
    wri.print();
    println!("Hello, world!");
}
