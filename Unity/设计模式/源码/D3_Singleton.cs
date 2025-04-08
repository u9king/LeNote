using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace DesignPattern3
{
    public class Singleton : MonoBehaviour
    {
        private static Singleton instance;

        public static Singleton Instance
        {
            get
            {
                if (instance == null)
                {
                    instance = new Singleton();
                }
                return instance;
            }
        }
        private Singleton() { }
    }
}
