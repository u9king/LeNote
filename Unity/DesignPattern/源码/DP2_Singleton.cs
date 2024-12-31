using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace DesignPattern
{
    public class Singleton_2 : MonoBehaviour
    {
        private static Singleton_2 instance = new Singleton_2();
        public static Singleton_2 Instance
        {
            get
            {
                if (instance == null)
                {
                    instance = new GameObject(nameof(Singleton_2)).AddComponent<Singleton_2>();
                }
                return instance;
            }
        }
    }
}
