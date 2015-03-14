//
//  ViewController.swift
//  BullsEye
//
//  Created by KevinQiangK on 3/13/15.
//  Copyright (c) 2015 KevinQiangK. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    var currentValue:Int = 50
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    @IBAction func showAlertMessage(){
        let message = "The valude of the slider is now:\(currentValue)"
        let alert = UIAlertController(title: "You Guess!", message: message, preferredStyle: .Alert)
        let action = UIAlertAction(title: "OK", style: .Default, handler: nil)
        alert.addAction(action)
        presentViewController(alert, animated: true, completion: nil)
    }

    @IBAction func slideMove(slider:UISlider){
        currentValue = lroundf(slider.value)
    }
}

