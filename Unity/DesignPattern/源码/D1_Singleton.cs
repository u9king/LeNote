using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace DesignPattern1
{
    public class Singleton
    {
        private static readonly Singleton instance = new Singleton();
        public static Singleton GetInstance()
        {
            return instance;
        }
        private Singleton() { }  //私有化构造函数
    }
}
