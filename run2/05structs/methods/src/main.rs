#[derive(Debug)]
struct Dog {
    name: String,
    weight: f32,
    height: f32,
}

impl Dog {
    fn get_name(&self) -> &str {
        &self.name[..]
    }
    // fn get_weight(&self) -> f32 {
    //     self.weight
    // }
    fn get_height(&self) -> f32 {
        self.height
    }
}

impl Dog {
    fn get_weight(&self) -> f32 {
        self.weight
    }
}

fn main() {
    let dog = Dog{
        name: String::from("ahei"),
        weight: 100.0,
        height: 75.5,
    };

    println!("dog: {:#?}", dog);

    println!("dog name: {}", dog.get_name());
    println!("dog weight: {}", dog.get_weight());
    println!("dog height: {}", dog.get_height());


    println!("Hello, world!");
}
