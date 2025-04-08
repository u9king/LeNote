using System;
using System.Collections;
using System.Collections.Generic;
using Unity.VisualScripting.YamlDotNet.Core.Tokens;
using UnityEngine;
namespace DesignPattern2
{
    public class Singleton
    {
        private static readonly Singleton instance;
       static Singleton()
        {
            instance = new Singleton();
        }

        public static Singleton GetInstance()
        {
            return instance;
        }
        private Singleton() { }
    }
}
