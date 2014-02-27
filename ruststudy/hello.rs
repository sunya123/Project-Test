fn main() {
	let hi = "hi";
	let mut count = 0;

	while count < 10 {
   		println!("count is {}", count);
    	count += 1;
	}

	println!("see {}",hi);

	println!("hello world");

	static MONSTER_FACTOR: f64 = 57.8;
	//let monster_size = MONSTER_FACTOR * 10.0;
	let monster_size: int = 50;
	println!("the data is {}",MONSTER_FACTOR);
	println(format!("repeat print the data {}",MONSTER_FACTOR));
	println(format!("the monster size is {}",monster_size));

	let price;
	let item="ddsdml";
	if item == "salad" {
    	price = 3.50;
	} else if item == "muffin" {
    	price = 2.25;
	} else {
    	price = 2.00;
	}
	println!("the price is {}",price)
	
}