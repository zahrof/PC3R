use std::sync::{Arc,Mutex,Condvar};
use std::collections::VecDeque;
use std::ops::Deref;

const PRODUITS: [&'static str; 12] = ["pomme", "poire","orange", "kiwi","banane",
                                        "abricot", "datte", "raisin", "pomelo",
                                        "grenade","litchi","mangue"];
const PROD_NB_THREAD: usize=12; //le nombre de threads doit être inférieur au nombre de produits
const CONS_NB_THREAD: usize=60;
const TAILLE_TAPIS: usize = 500;
const CIBLE_PRODUCTION: usize =5;

struct Paquet {
    nom_produit : String
}

struct Tapis {
    produits : VecDeque<Paquet>,
    taille : usize
}

fn defile(t: &Arc<(Mutex<Tapis>,Condvar)>)-> Paquet{
    let (lgfi,cv)=t.deref();

    let mut fifo = lgfi.lock().unwrap();
    while fifo.produits.is_empty(){
        fifo=cv.wait(fifo).unwrap();
    }
    if fifo.produits.capacity()==fifo.taille {
        cv.notify_all();
    }

    let p = fifo.produits.pop_front().unwrap();
    p

}

fn empile(t: Arc<(Mutex<Tapis>,Condvar)>,p: Paquet) -> Arc<(Mutex<Tapis>,Condvar)>{

    let  tapis_clone = Arc::clone(&t);

    let (lgfi,cv)=&*t;

    let mut fifo=lgfi.lock().unwrap();


    while fifo.produits.capacity()==fifo.taille {
        fifo=cv.wait(fifo).unwrap();
    }
    if fifo.produits.is_empty() {
        cv.notify_all();
    }
    fifo.taille +=1;
    fifo.produits.push_back(p);
   tapis_clone
}
fn main(){
    let fi = Tapis {produits : VecDeque::new(),taille : TAILLE_TAPIS};
    let finis = Arc::new(Mutex::<usize>::new(PROD_NB_THREAD * CIBLE_PRODUCTION));
    let tapis = Arc::new((Mutex::new(fi),Condvar::new()));
    let mut handles = Vec::new();

    for i in 0..PROD_NB_THREAD {
        let mut  tapis_clone = Arc::clone(&tapis);
        handles.push(std::thread::spawn(move || {
            for j in 0..CIBLE_PRODUCTION {
                let p = Paquet {nom_produit : format!{"{}{}",PRODUITS[i],j.to_string()} };
                println!("Producteur {} écrit {}", i.to_string(),p.nom_produit);
                tapis_clone=empile(tapis_clone,p);
            }

        }));
    };

    for i in 0..CONS_NB_THREAD{
        let tapis_clone = Arc::clone(&tapis);
        let finis_clone = Arc::clone(&finis);
        handles.push(std::thread::spawn(move || {
            let mut continu = true;
            while continu {
                {
                    let mut f_l = finis_clone.lock().unwrap();
                    if *f_l <= 0 {continu = false;} else {*f_l -= 1;
                       println!("Reste {}.", f_l.to_string())}
                };
                if continu {
                    let p = defile(&tapis_clone);
                    println!("C{} mange {}", i.to_string(), p.nom_produit.to_string());
                }
            };

        }));
    };

    for h in handles {
        h.join().unwrap();
    }
}