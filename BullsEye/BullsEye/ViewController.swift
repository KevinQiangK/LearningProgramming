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
    
    var _totalScores = 0
    var _totalRounds = 0
    
    @IBOutlet weak var slider:UISlider!
    @IBOutlet weak var targetLabel:UILabel!
    @IBOutlet weak var totalScoresLabel:UILabel!
    @IBOutlet weak var totalRoundsLabel:UILabel!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        startNewGame()
        updateLabels()
    }

    func startNewGame(){
        _totalScores = 0
        _totalRounds = 0
        startNewRound()
    }
    
    func startNewRound(){
        targetValue = 1 + Int(arc4random_uniform(100))
        currentValue = lroundf(slider.value)
        slider.value = Float(currentValue)
        _totalRounds += 1
    }
    
    func updateLabels(){
        targetLabel.text = String(targetValue)
        totalScoresLabel.text = String(_totalScores)
        totalRoundsLabel.text = String(_totalRounds)
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    @IBAction func showAlertMessage(){
        let difference = abs(targetValue - currentValue)
        var point = 100 - difference
        
        var title:String
        if(difference == 0){
            title = "Pefect"
            point += 100
        }else if(difference < 5){
            title = "Great"
            point += 50
        }else{
            title = "Not very closed to."
        }
        
        _totalScores += point
        let message = "You get \(point) points."
        let alert = UIAlertController(title: "You Guess!", message: message, preferredStyle: .Alert)
        let action = UIAlertAction(title: title, style: .Default,
            handler: { action in self.startNewRound()
        self.updateLabels()})
        
        alert.addAction(action)
        presentViewController(alert, animated: true, completion: nil)
    }

    @IBAction func slideMove(slider:UISlider){
        currentValue = lroundf(slider.value)
    }
    
    @IBAction func startOverGame(){
        startNewGame()
        updateLabels()
    }
}

