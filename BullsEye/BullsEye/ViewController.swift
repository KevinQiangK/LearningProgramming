//
//  ViewController.swift
//  BullsEye
//
//  Created by KevinQiangK on 3/13/15.
//  Copyright (c) 2015 KevinQiangK. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    var currentValue:Int = 0
    var targetValue:Int = 0
    
    @IBOutlet weak var slider:UISlider!
    @IBOutlet weak var targetLabel:UILabel!
    override func viewDidLoad() {
        super.viewDidLoad()
        startNewRound()
        updateLabels()
    }

    func startNewRound(){
        targetValue = 1 + Int(arc4random_uniform(100))
        currentValue = lroundf(slider.value)
        slider.value = Float(currentValue)
    }
    
    func updateLabels(){
        targetLabel.text = String(targetValue)
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    @IBAction func showAlertMessage(){
        let message = "The current value is: \(currentValue)"
                    + "\nThe target value is: \(targetValue)";
        
        let alert = UIAlertController(title: "You Guess!", message: message, preferredStyle: .Alert)
        let action = UIAlertAction(title: "OK", style: .Default, handler: nil)
        alert.addAction(action)
        presentViewController(alert, animated: true, completion: nil)
        startNewRound()
        updateLabels()
    }

    @IBAction func slideMove(slider:UISlider){
        currentValue = lroundf(slider.value)
    }
}

