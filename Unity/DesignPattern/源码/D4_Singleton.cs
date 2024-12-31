using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace DesignPattern4
{
    public class Singleton
    {
        private static Singleton instance;

        private static readonly object lockObj = new object();  //¶àÏß³ÌËø

        public static Singleton Instance
        {
            get
            {
                if (instance == null)
                {
                    lock (lockObj)
                    {
                        if (instance == null)
                        {
                            instance = new Singleton();
                        }
                    }
                }
                return instance;
            }
        }
        private Singleton() { }
    }
}
