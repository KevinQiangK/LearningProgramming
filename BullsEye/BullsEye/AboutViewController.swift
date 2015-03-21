//
//  AboutViewController.swift
//  BullsEye
//
//  Created by KevinQiangK on 3/19/15.
//  Copyright (c) 2015 KevinQiangK. All rights reserved.
//

import UIKit

class AboutViewController: UIViewController {
    
    @IBOutlet weak var webView:UIWebView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        LoadAboutHtml()
    }
    
    @IBAction func closeAboutView(){
        dismissViewControllerAnimated(true, completion:nil)
    }
    
    func LoadAboutHtml(){
    
        if let htmlFile = NSBundle.mainBundle().pathForResource("BullsEye", ofType: "html"){
            let htmlData = NSData(contentsOfFile: htmlFile)
            let baseUrl = NSURL.fileURLWithPath(NSBundle.mainBundle().bundlePath)
            webView.loadData(htmlData, MIMEType: "text/html", textEncodingName: "UTF-8", baseURL: baseUrl)
        }
    }

}
